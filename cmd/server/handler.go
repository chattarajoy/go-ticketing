package server

import (
	"fmt"
	"net/http"
)

func defaultHandler(handlerType string) http.Handler {
	if handlerType == "healthz" {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			healthCheckStatus(request, writer)
		})
	}
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		NotFoundHandler(request, writer)
	})
}

func healthCheckStatus(request *http.Request, writer http.ResponseWriter) {
	var healthResponse = `
				{
					"name": "apiServer",
					"status": "OK"
        }`
	writer.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(writer, healthResponse)
}

func NotFoundHandler(request *http.Request, writer http.ResponseWriter) {
	var healthResponse = `
				{
					"name": "NotFound",
					"status": "404"
				}`
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(writer, healthResponse)
}
