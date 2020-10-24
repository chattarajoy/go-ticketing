package router_test

// This file includes tests for all types of routers defined in this package
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"commerceiq.ai/ticketing/internal/router"
)

func dummyHandler(status string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var htmlIndex = `{ "name": "dummy", "status": "` + status + `" }`
		_, _ = fmt.Fprintf(writer, htmlIndex)
	})
}

func TestRouters_Handle(t *testing.T) {

	tests := []struct {
		name     string
		router   router.Router
		path     string
		expected string
		status   int
	}{
		{"HTTP Router 200", router.CreateRouter("httprouter"), "/dummy",
			`{ "name": "dummy", "status": "HTTP Router 200" }`, http.StatusOK},
		{"HTTP Router 404", router.CreateRouter("httprouter"), "/random", "", http.StatusNotFound},
		{"NET/HTTP Router 200", router.CreateRouter("nethttp"), "/dummy",
			`{ "name": "dummy", "status": "NET/HTTP Router 200" }`, http.StatusOK},
		{"NET/HTTP Router 404", router.CreateRouter("nethttp"), "/random", "", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpRouter := tt.router
			httpRouter.Handle("GET", tt.path, dummyHandler(tt.name))
			req, err := http.NewRequest("GET", "/dummy", nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpRouter.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}

			if tt.expected != "" {
				expected := tt.expected
				if rr.Body.String() != expected {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), expected)
				}
			}
		})
	}
}

func TestRouters_Name(t *testing.T) {
	tests := []struct {
		name   string
		router router.Router
	}{
		{"HTTP Router", router.CreateRouter("httprouter")},
		{"Net Http Router", router.CreateRouter("nethttp")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routeHandler := tt.router
			if got := routeHandler.Name(); got != tt.name {
				t.Errorf("Name() = %v, want %v", got, tt.name)
			}
		})
	}
}

func TestRouters_NotFound(t *testing.T) {

	tests := []struct {
		name            string
		path            string
		expected        string
		router          router.Router
		handler         http.Handler
		notFoundHandler http.Handler
	}{
		{"HTTP - Registered Path", "/test", `{ "name": "dummy", "status": "http/ok" }`,
			router.CreateRouter("httprouter"), dummyHandler("http/ok"), dummyHandler("http/NotFound")},
		{"HTTP - Not Registered Path", "/xyz", `{ "name": "dummy", "status": "http/NotFound" }`,
			router.CreateRouter("httprouter"), dummyHandler("http/ok"), dummyHandler("http/NotFound")},
		{"NetHTTP - Registered Path", "/test", `{ "name": "dummy", "status": "nethttp/ok" }`,
			router.CreateRouter("nethttp"), dummyHandler("nethttp/ok"), dummyHandler("nethttp/NotFound")},
	}
	for _, tt := range tests {

		httpRouter := tt.router
		httpRouter.Handle("GET", "/test", tt.handler)
		httpRouter.NotFound(tt.notFoundHandler)

		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			httpRouter.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := tt.expected
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}

func TestCreateRouter(t *testing.T) {
	testRouter := router.CreateRouter("random")
	if testRouter != nil {
		t.Fatalf("Default router should be nil, got object")
	}
}
