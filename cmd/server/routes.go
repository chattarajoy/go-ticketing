package server

import (
	"commerceiq.ai/ticketing/internal/server"
)

func (s *Server) setupRoutes() {
	s.Routes = []server.Route{
		{
			Method:  "GET",
			Path:    "/healthz",
			Handler: defaultHandler("healthz"),
		},
	}
}
