package api

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"

	"commerceiq.ai/ticketing/internal/cache"
	jsonHelper "commerceiq.ai/ticketing/internal/json"
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

type Contract interface {
	Validate(db *gorm.DB) error
}

func ValidateContract(c Contract, request *http.Request, writer http.ResponseWriter, db *gorm.DB) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(c)
	if err == nil {
		err = c.Validate(db)
	}
	if err != nil {
		resp := ErrorResponse(err.Error(), http.StatusBadRequest)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
		return err
	}
	return nil
}
