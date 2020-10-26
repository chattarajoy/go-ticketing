package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type statusWriter struct {
	http.ResponseWriter
	status    int
	length    int
	requestID string
}

func (writer *statusWriter) Write(b []byte) (int, error) {
	if writer.status == 0 {
		writer.status = http.StatusOK
	}

	length, err := writer.ResponseWriter.Write(b)
	writer.length += length
	return length, err

}

func (hServer *HttpServer) wrapHandler(handler http.Handler) http.Handler {
	next := hServer.LogHandler(hServer.AccessControlHandler(handler))
	for _, nextHandler := range hServer.Input.WrapHandlers {
		next = nextHandler(next)
	}
	return hServer.RecoverHandler(next)
}

func (hServer *HttpServer) RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				reqId := uuid.New().String()
				jsonBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprintf("Something went wrong. Error ID: %s", reqId),
				})
				_ = hServer.Logger.Log("Request-Id", reqId, "Error: ", err, "requestUrl", request.URL)
				writer.Header().Set("X-Request-Id", reqId)
				writer.Header().Set("Content-Type", "application/json")
				_, _ = writer.Write(jsonBody)
			}
		}()
		next.ServeHTTP(writer, request)
	})
}

func (hServer *HttpServer) LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		statWriter := statusWriter{ResponseWriter: writer}

		defer func(begin time.Time, request *http.Request) {
			_ = hServer.Logger.Log("Request-Id", statWriter.requestID, "host", request.Host, "path", request.URL.Path, "remote",
				request.RemoteAddr, "method", request.Method, "status", statWriter.status,
				"content-length", statWriter.length, "took", time.Since(begin))
		}(time.Now(), request)

		next.ServeHTTP(&statWriter, request)
	})
}

func (hServer *HttpServer) AccessControlHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		writer.Header().Set("X-Request-Id", uuid.New().String())
		if request.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(writer, request)
	})
}
