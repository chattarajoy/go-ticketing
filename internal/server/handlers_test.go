package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"commerceiq.ai/ticketing/internal/server"
	"commerceiq.ai/ticketing/internal/testhelpers"
)

func Test_httpServer_accessControlHandler(t *testing.T) {
	type fields struct {
		ServerInput *server.ServerInput
	}
	type args struct {
		next   http.Handler
		path   string
		status int
		result string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Test Headers", fields{GetServerInput(false)},
			args{DummyHandler("ok"), "/test", http.StatusOK, `{ "name": "dummy", "status": "ok" }`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hServer := &server.HttpServer{
				ServerInput: tt.fields.ServerInput,
			}
			hServer.Router.Handle("GET", tt.args.path, hServer.AccessControlHandler(tt.args.next))
			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			hServer.Router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.args.status)
			}

			if tt.args.result != "" {
				expected := tt.args.result
				if rr.Body.String() != expected {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), expected)
				}
			}
			headers := rr.Header()
			expectedHeaders := make(map[string][]string)
			expectedHeaders["Access-Control-Allow-Origin"] = append(expectedHeaders["Access-Control-Allow-Origin"], "*")
			expectedHeaders["Access-Control-Allow-Methods"] = append(expectedHeaders["Access-Control-Allow-Methods"], "GET, POST, PUT, DELETE, PATCH, OPTIONS")
			expectedHeaders["Access-Control-Allow-Headers"] = append(expectedHeaders["Access-Control-Allow-Headers"], "Origin, Content-Type")
			expectedHeaders["X-Request-Id"] = append(expectedHeaders["X-Request-Id"], "-")
			for key, value := range expectedHeaders {
				if strings.Contains(headers[key][0], value[0]) != true {
					t.Errorf("Expected %s to equal %s got %s", key, value[0], headers[key][0])
				}
			}
		})
	}
}

func Test_httpServer_logHandler(t *testing.T) {
	type fields struct {
		ServerInput *server.ServerInput
	}
	type args struct {
		next   http.Handler
		path   string
		status int
		result string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Test Log Output", fields{GetServerInput(false)},
			args{DummyHandler("ok"), "/test", http.StatusOK, `{ "name": "dummy", "status": "ok" }`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hServer := &server.HttpServer{
				ServerInput: tt.fields.ServerInput,
			}
			var writer *bytes.Buffer
			hServer.Logger, writer = testhelpers.LoggerWithWriter()
			hServer.Router.Handle("GET", tt.args.path, hServer.LogHandler(tt.args.next))
			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			hServer.Router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.args.status)
			}

			if tt.args.result != "" {
				expected := tt.args.result
				if rr.Body.String() != expected {
					t.Errorf("handler returned unexpected body: got %v want %v",
						rr.Body.String(), expected)
				}
			}
			outLog := writer.String()
			expectedConatins := []string{
				"Request-Id=", "host=", "path=/test", "remote=", "method=GET", "status=200", "content-length=35", "took=",
			}
			for _, str := range expectedConatins {
				if strings.Contains(outLog, str) != true {
					t.Errorf("Expected log to contain %s, not found. Log: %s", str, outLog)
				}
			}
		})
	}
}

func panicHandler(status string) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		panic("I need to panic!")
	})
}

func Test_httpServer_recoverHandler(t *testing.T) {
	type fields struct {
		ServerInput *server.ServerInput
	}
	type args struct {
		next   http.Handler
		path   string
		status int
		result string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Test Recovery", fields{GetServerInput(false)},
			args{panicHandler("ok"), "/test", http.StatusInternalServerError,
				`{"error":"Something went wrong. Error ID: `}},
		{"Test Without Panic", fields{GetServerInput(false)},
			args{DummyHandler("ok"), "/test", http.StatusOK,
				`{ "name": "dummy", "status": "ok" }`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hServer := &server.HttpServer{
				ServerInput: tt.fields.ServerInput,
			}
			hServer.Router.Handle("GET", tt.args.path, hServer.RecoverHandler(tt.args.next))
			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatalf(err.Error())
			}

			rr := httptest.NewRecorder()
			hServer.Router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.args.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.args.status)
			}

			if tt.args.result != "" {
				expected := tt.args.result
				if strings.Contains(rr.Body.String(), expected) != true {
					t.Errorf("handler returned unexpected body: got %v want like %v",
						rr.Body.String(), expected)
				}
			}
		})
	}
}

func customHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Custom Header", "custom")
		next.ServeHTTP(writer, request)
	})
}

func Test_wrapHandlers(t *testing.T) {

	type fields struct {
		ServerInput *server.ServerInput
	}
	tests := []struct {
		name            string
		fields          fields
		path            string
		status          int
		result          string
		expectedHeaders map[string][]string
	}{
		{"basic routing", fields{GetServerInput(false)},
			"/test", http.StatusOK, `{ "name": "dummy", "status": "test" }`,
			map[string][]string{
				"Custom Header": {"custom"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hServer := &server.HttpServer{
				ServerInput: tt.fields.ServerInput,
			}
			hServer.ServerInput.WrapHandlers = []func(next http.Handler) http.Handler{customHeaderHandler}
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

			headers := rr.Header()
			for key, val := range tt.expectedHeaders {
				if headers[key][0] != val[0] {
					t.Errorf("Expected header %s to be %s got %s", key, val, headers[key][0])
				}
			}
		})
	}
}
