package api

import (
	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	"commerceiq.ai/ticketing/pkgs/service/cinema"
)

type Handler struct {
	db *gorm.DB

	// services
	svc *HandlerServices
}

type HandlerServices struct {
	cinema *cinema.Service
}

func NewAPIHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
		svc: &HandlerServices{
			cinema: cinema.NewService(db,
				cache.NewCache(cache.InMemoryCache))},
	}
}
