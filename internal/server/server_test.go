package server_test

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"syscall"
	"testing"
	"time"

	"commerceiq.ai/ticketing/internal/workgroup"

	"commerceiq.ai/ticketing/internal/router"
	"commerceiq.ai/ticketing/internal/server"
	"commerceiq.ai/ticketing/internal/testhelpers"
)

func DummyHandler(status string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var htmlIndex = `{ "name": "dummy", "status": "` + status + `" }`
		_, _ = fmt.Fprintf(writer, htmlIndex)
	})
}

func GetServerInput(withNotFoundHandler bool) *server.ServerInput {

	max, min := 60000, 30000
	logger := testhelpers.FakeLogger()
	serverInput := &server.ServerInput{
		Port:   rand.Intn(max-min) + min,
		Router: router.CreateRouter("httprouter"),
		Logger: logger,
		RouteMap: []server.Route{
			{"GET", "/test", DummyHandler("test")},
			{"GET", "/test2", DummyHandler("test2")},
		},
		ServerDrainTime: 1,
	}
	if withNotFoundHandler == true {
		serverInput.NotFoundHandler = DummyHandler("NotFound")
	}
	return serverInput
}

func TestCreateServer(t *testing.T) {
	t.Parallel()
	var group workgroup.Group

	type args struct {
		group *workgroup.Group
		inp   *server.ServerInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"RunServer", args{&group, GetServerInput(false)}, false},
	}

	stop := make(chan int, 1)

	for _, tt := range tests {
		t.Run(tt.name+"Create Server", func(t *testing.T) {
			t.Parallel()
			if err := server.CreateServer(tt.args.group, tt.args.inp); (err != nil) != tt.wantErr {
				t.Errorf("CreateServer() error = %v, wantErr %v", err, tt.wantErr)
			}

			_ = tt.args.group.Run()
			stop <- 0

		})

		t.Run(tt.name+"Kill Server", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 1)
			// send sigint to kill
			_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)

			select {
			case <-stop:
				t.Log("Server Stopped!")
			case <-time.After(5 * time.Second):
				t.Fatalf("Server did not shut down")
			}
		})
	}
}

func Test_httpServer_routes(t *testing.T) {

	type fields struct {
		ServerInput *server.ServerInput
	}
	tests := []struct {
		name   string
		fields fields
		path   string
		status int
		result string
	}{
		{"basic routing", fields{GetServerInput(false)},
			"/test", http.StatusOK, `{ "name": "dummy", "status": "test" }`},
		{"routing 2", fields{GetServerInput(false)},
			"/test2", http.StatusOK, `{ "name": "dummy", "status": "test2" }`},
		{"routing NotFound", fields{GetServerInput(true)},
			"/random", http.StatusOK, `{ "name": "dummy", "status": "NotFound" }`},
		{"routing NotFound 404", fields{GetServerInput(false)},
			"/random", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hServer := &server.HttpServer{
				ServerInput: tt.fields.ServerInput,
			}
			hServer.Routes()
			httpRouter := hServer.Router
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpRouter.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}

			if tt.result != "" {
				expected := tt.result
				if rr.Body.String() != expected {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), expected)
				}
			}
		})
	}
}

func Test_httpServer_shutDownServer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		*server.ServerInput
	}{
		{"testShutdown", GetServerInput(false)},
	}
	for _, tt := range tests {

		hServer := &server.HttpServer{
			ServerInput: tt.ServerInput,
		}
		server := &http.Server{Addr: ":" + fmt.Sprint(hServer.Port), Handler: hServer.Router}
		stopAll := make(chan int, 1)
		stop := make(chan struct{})

		t.Run(tt.name+"run shutDown Server", func(t *testing.T) {
			t.Parallel()
			hServer.ShutDownServer(server, stop)
		})

		t.Run(tt.name+"run Server ListenAndServe", func(t *testing.T) {
			t.Parallel()
			_ = server.ListenAndServe()
			stopAll <- 0
		})

		t.Run(tt.name+"shut down server and wait", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second * 1)
			stop <- struct{}{}
			select {
			case <-stopAll:
				t.Log("Server Stopped!")
			case <-time.After(5 * time.Second):
				t.Fatalf("Server did not shut down")
			}
		})
	}
}
