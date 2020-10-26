package api

import (
	"net/http"
)

type Response struct {
	Success      bool
	StatusCode   int
	Data         []byte
	ErrorMessage string
}

func ErrorResponse(errorMessage string, statusCode int) *Response {
	return &Response{
		Success:      false,
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
}

func SuccessResponse(payload []byte) *Response {
	return &Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       payload,
	}
}
