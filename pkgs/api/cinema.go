package api

import (
	"encoding/json"
	"net/http"

	jsonHelper "commerceiq.ai/ticketing/internal/json"
	"commerceiq.ai/ticketing/pkgs/service/cinema"
)

func (h *Handler) ListCinemas(request *http.Request, writer http.ResponseWriter) {
	if out, err := h.svc.cinema.ListCinemas(); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddCinema(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp cinema.AddCinemaInput
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&inp)
	if err == nil {
		err = inp.Validate(h.db)
	}
	if err != nil {
		resp := ErrorResponse(err.Error(), http.StatusBadRequest)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
		return
	}

	// process input and return output
	if out, err := h.svc.cinema.AddCinema(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) AddCinemaScreen(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp cinema.AddCinemaScreenInput
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&inp)
	if err == nil {
		err = inp.Validate(h.db)
	}
	if err != nil {
		resp := ErrorResponse(err.Error(), http.StatusBadRequest)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
		return
	}

	// process input and return output
	if out, err := h.svc.cinema.AddCinemaScreen(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}
