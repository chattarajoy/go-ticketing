package server

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/chattarajoy/go-ticketing/internal/router"
	"github.com/chattarajoy/go-ticketing/internal/server"
)

func TestServer_BasicRoutes(t *testing.T) {
	type fields struct {
		Config *Config
		Logger log.Logger
		Name   string
		Router router.Router
		Routes []server.Route
	}
	type args struct {
		handler http.Handler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "healthz",
			args: args{
				handler: defaultHandler(""),
			},
			wantErr: false,
			fields: fields{
				Config: &Config{
					HTTPPort:        3838,
					ServerDrainTime: 5,
					Debug:           0,
				},
				Logger: log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)),
				Name:   "APIServer",
				Router: router.CreateRouter("httprouter"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"_Server", func(t *testing.T) {
			t.Parallel()
			s := &Server{
				Config: tt.fields.Config,
				Logger: tt.fields.Logger,
				Name:   tt.fields.Name,
				Router: tt.fields.Router,
				Routes: tt.fields.Routes,
			}
			s.setupRoutes()
			if err := s.runServer(tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("runServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Run(tt.name+"_Shutdown", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 3)
			defer syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		})

		t.Run(tt.name+"_Request", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 2)
			res, err := http.Get("http://localhost:3838/healthz")
			if err != nil {
				t.Error(err)
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("http request error wanted status 200, got: %d", res.StatusCode)
			}
		})

		t.Run(tt.name+"_Request_404", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 2)
			res, err := http.Get("http://localhost:3838/random")
			if err != nil {
				t.Error(err)
			}
			if res.StatusCode != http.StatusNotFound {
				t.Errorf("http request error wanted status 200, got: %d", res.StatusCode)
			}
		})
	}
}
