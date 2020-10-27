package api

import (
	"net/http"

	jsonHelper "commerceiq.ai/ticketing/internal/json"
	"commerceiq.ai/ticketing/pkgs/service/booking"
)

func (h *Handler) ListBookings(request *http.Request, writer http.ResponseWriter) {
	if out, err := h.svc.booking.ListBookings(); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}

func (h *Handler) BookSeats(request *http.Request, writer http.ResponseWriter) {
	// validate input
	var inp booking.BookSeatsInput
	if err := ValidateContract(&inp, request, writer, h.db); err != nil {
		return
	}

	// process input and return output
	if out, err := h.svc.booking.BookSeats(&inp); err != nil {
		resp := ErrorResponse(err.Error(), http.StatusUnprocessableEntity)
		jsonHelper.WriteResult(&resp, writer, http.StatusUnprocessableEntity)
	} else {
		resp := SuccessResponse(out)
		jsonHelper.WriteResult(&resp, writer, http.StatusOK)
	}
}
