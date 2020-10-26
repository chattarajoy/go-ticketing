package server

import (
	"commerceiq.ai/ticketing/internal/server"
	"commerceiq.ai/ticketing/pkgs/api"
)

func (s *Server) setupRoutes() {
	apiHandler := api.NewAPIHandler(s.Db)
	s.Routes = []server.Route{
		{
			Method:  "GET",
			Path:    "/healthz",
			Handler: defaultHandler("healthz"),
		},
		{
			Method:  "GET",
			Path:    "/cinemas",
			Handler: wrapHandler(apiHandler.ListCinemas),
		},
		{
			Method:  "POST",
			Path:    "/cinema",
			Handler: wrapHandler(apiHandler.AddCinema),
		},
		{
			Method:  "POST",
			Path:    "/screen",
			Handler: wrapHandler(apiHandler.AddCinemaScreen),
		},
	}
}
