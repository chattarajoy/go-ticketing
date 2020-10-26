package server

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"

	"commerceiq.ai/ticketing/internal/router"
	"commerceiq.ai/ticketing/internal/server"
	"commerceiq.ai/ticketing/internal/workgroup"
)

type CMD struct {
	RootCmd *cobra.Command
	Logger  log.Logger
}

type Server struct {
	Config *Config
	Logger log.Logger
	Name   string
	Router router.Router
	Routes []server.Route
}

type Config struct {
	HTTPPort        int     `json:"http_port" usage:"server http port number"`
	HTTPSPort       int     `json:"https_port" usage:"server https port number"`
	ServerDrainTime int     `json:"server_drain_time" usage:"number of seconds needed for server to shutdown"`
	Debug           float64 `json:"debug" usage:"enable pprof to debug" required:"true"`
}

var (
	apiServer = &Server{
		Config: &Config{
			HTTPPort:        4000,
			ServerDrainTime: 2,
		},
		Name:   "APIServer",
		Router: router.CreateRouter("httprouter"),
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start API Server",
		Long:  "",
		Run:   startServer,
	}
)

func Init(cmd *CMD) {
	cmd.RootCmd.AddCommand(serverCmd)
	apiServer.initLogger(cmd.Logger)
	// cfg, err := apiServer.loadConfig(serverCmd.Flags())
	// if err != nil {
	// 	_ = apiServer.Logger.Log("Error Reading Config: ", err, "Exiting", "!!")
	// 	os.Exit(1)
	// }
	// _ = cfg.Print(apiServer.Logger)
	apiServer.setupRoutes()
}

func (s *Server) initLogger(logger log.Logger) {
	s.Logger = log.With(logger, "app", s.Name)
}

func startServer(_ *cobra.Command, _ []string) {
	err := apiServer.runServer(defaultHandler(""))
	if err != nil {
		_ = apiServer.Logger.Log("Error", err.Error(), "Exiting", "...")
		os.Exit(1)
	}
}

func (s *Server) runServer(handler http.Handler) error {
	var group workgroup.Group

	err := server.CreateServer(&group, &server.Input{
		Port:            s.Config.HTTPPort,
		Router:          s.Router,
		Logger:          s.Logger,
		NotFoundHandler: handler,
		ServerDrainTime: s.Config.ServerDrainTime,
		Routes:          s.Routes,
	})

	if err != nil {
		return err
	}

	_ = s.Logger.Log("Start", "Server", "App", s.Name)
	_ = s.Logger.Log("WorkGroup", "Shutdown", "Status", group.Run())
	return nil
}
