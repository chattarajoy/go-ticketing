package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/chattarajoy/go-ticketing/internal/router"
	"github.com/chattarajoy/go-ticketing/internal/server"
	"github.com/chattarajoy/go-ticketing/internal/workgroup"
	"github.com/chattarajoy/go-ticketing/pkgs/models"
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
	Db     *gorm.DB
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
}

func (s *Server) initLogger(logger log.Logger) {
	s.Logger = log.With(logger, "app", s.Name)
}

func startServer(_ *cobra.Command, _ []string) {
	err := apiServer.dbInit()
	if err != nil {
		_ = apiServer.Logger.Log("Error", err.Error(), "Exiting", "...")
		os.Exit(1)
	}

	apiServer.setupRoutes()
	err = apiServer.runServer(defaultHandler(""))
	if err != nil {
		_ = apiServer.Logger.Log("Error", err.Error(), "Exiting", "...")
		os.Exit(1)
	}
}

func (s *Server) dbInit() error {
	// TODO: Move to CLI Arguments
	dsn := "user:user@tcp(db:3306)/ticketing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to DB: ", err.Error())
		os.Exit(1)
	}
	s.Db = db
	return db.AutoMigrate(&models.Booking{},
		&models.BookingSeat{},
		&models.Movie{},
		&models.MovieShow{},
		&models.User{},
		&models.City{},
		&models.Cinema{},
		&models.CinemaScreen{},
		&models.CinemaSeat{},
	)

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
