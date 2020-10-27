package server

import (
	"github.com/chattarajoy/go-ticketing/internal/server"
	"github.com/chattarajoy/go-ticketing/pkgs/api"
)

func (s *Server) setupRoutes() {
	apiHandler := api.NewAPIHandler(s.Db)
	s.Routes = []server.Route{
		{
			Method:  "GET",
			Path:    "/healthz",
			Handler: defaultHandler("healthz"),
		},
		// cinema routes
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
		// movie routes
		{
			Method:  "GET",
			Path:    "/shows",
			Handler: wrapHandler(apiHandler.GetShow),
		},
		{
			Method:  "POST",
			Path:    "/show",
			Handler: wrapHandler(apiHandler.AddShow),
		},
		{
			Method:  "POST",
			Path:    "/movie",
			Handler: wrapHandler(apiHandler.AddMovie),
		},
		// booking related routes
		{
			Method:  "GET",
			Path:    "/bookings",
			Handler: wrapHandler(apiHandler.ListBookings),
		},
		{
			Method:  "POST",
			Path:    "/book",
			Handler: wrapHandler(apiHandler.BookSeats),
		},
	}
}
