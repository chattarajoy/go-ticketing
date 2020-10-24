package router

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type HTTPRouter struct {
	handler http.Handler
	router  func() *httprouter.Router
}

func NewHTTPRouter(handler http.Handler) *HTTPRouter {
	httpRouter := &HTTPRouter{handler: handler}
	httpRouter.router = func() *httprouter.Router {
		return httpRouter.handler.(*httprouter.Router)
	}
	return httpRouter
}

func (httpRouter *HTTPRouter) Name() string {
	return "HTTP Router"
}

func (httpRouter *HTTPRouter) Handle(method, path string, handler http.Handler) {
	allowedMethods := []string{"head", "get", "post", "put", "patch", "delete", "options"}
	curMethod := strings.ToLower(method)
	for _, allowedMethod := range allowedMethods {
		if allowedMethod == curMethod {
			httpRouter.router().Handler(method, path, handler)
			return
		}
	}
}

func (httpRouter *HTTPRouter) NotFound(handler http.Handler) {
	httpRouter.router().NotFound = handler
}

func (httpRouter *HTTPRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	httpRouter.handler.ServeHTTP(writer, request)
}
