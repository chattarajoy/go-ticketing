package router

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Router interface {
	NotFound(handler http.Handler)
	http.Handler
	Handle(method, path string, handler http.Handler)
	Name() string
}

func CreateRouter(routerName string) Router {
	switch strings.ToLower(routerName) {
	case "httprouter":
		return NewHTTPRouter(httprouter.New())
	case "nethttp":
		return NewNetHTTP(http.NewServeMux())
	default:
		return nil
	}
}
