package api

import (
	"net/http"

	jsonHelper "github.com/chattarajoy/go-ticketing/internal/json"
	"github.com/chattarajoy/go-ticketing/pkgs/service/movie"
)

func (h *Handler) GetShow(request *http.Request, writer http.ResponseWriter) {
	var inp movie.GetMovieShowInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	if out, err := h.svc.movie.GetMovieShow(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddMovie(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp movie.AddMovieInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	// process input and return output
	if out, err := h.svc.movie.AddMovie(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddShow(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp movie.AddMovieShowInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	// process input and return output
	if out, err := h.svc.movie.AddMovieShow(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}
