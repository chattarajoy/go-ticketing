package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/chattarajoy/go-ticketing/internal/router"
	"github.com/chattarajoy/go-ticketing/internal/workgroup"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

type Input struct {
	Port            int
	Router          router.Router
	Logger          log.Logger
	NotFoundHandler http.Handler
	ServerDrainTime int
	Routes          []Route
	WrapHandlers    []func(next http.Handler) http.Handler
}

type HttpServer struct {
	*Input
}

func CreateServer(group *workgroup.Group, inp *Input) error {
	hSever := &HttpServer{Input: inp}
	hSever.SetupRoutes()
	server := &http.Server{Addr: ":" + fmt.Sprint(inp.Port), Handler: hSever.Router}

	group.Add(func(stop <-chan struct{}) error {
		go hSever.ShutDownServer(server, stop)
		err := server.ListenAndServe()
		if err != nil {
			_ = hSever.Logger.Log("Error", "ListenAndServe", "Error", err.Error())
		}
		return err
	})

	group.Add(func(stop <-chan struct{}) error {
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)

		<-sigChannel
		return fmt.Errorf("process Interupted")
	})

	return nil
}

func (hServer *HttpServer) ShutDownServer(server *http.Server, stop <-chan struct{}) {
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(hServer.ServerDrainTime)*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		_ = hServer.Logger.Log("Shutdown", "Failed", "For", server.Addr, "Error", err.Error())
	} else {
		_ = hServer.Logger.Log("Shutdown", "Success", "For", server.Addr)
	}
}

func (hServer *HttpServer) SetupRoutes() {
	// setup routes from route map
	for _, route := range hServer.Routes {
		hServer.handle(route.Method, route.Path, route.Handler)
	}

	// setup not found handler if defined
	if hServer.NotFoundHandler != nil {
		hServer.Router.NotFound(hServer.wrapHandler(hServer.NotFoundHandler))
	}
}

func (hServer *HttpServer) handle(method, path string, handler http.Handler) {
	hServer.Router.Handle(method, path, hServer.wrapHandler(handler))
}
