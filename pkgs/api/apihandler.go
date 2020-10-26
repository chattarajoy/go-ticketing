package api

import (
	"net/http"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	"commerceiq.ai/ticketing/pkgs/service/booking"
	"commerceiq.ai/ticketing/pkgs/service/cinema"
	"commerceiq.ai/ticketing/pkgs/service/movie"
)

// TODO: Add Logger
type Handler struct {
	db *gorm.DB

	// services
	svc *HandlerServices
}

type HandlerServices struct {
	cinema  *cinema.Service
	movie   *movie.Service
	booking *booking.Service
}

func NewAPIHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
		svc: &HandlerServices{
			cinema: cinema.NewService(db,
				cache.NewCache(cache.InMemoryCache)),
			movie: movie.NewService(db,
				cache.NewCache(cache.InMemoryCache)),
			booking: booking.NewService(db,
				cache.NewCache(cache.InMemoryCache)),
		},
	}
}

type HandlerFunc func(request *http.Request, writer http.ResponseWriter)
