package api

import (
	"net/http"
)

type Response struct {
	Success      bool        `json:"success"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

func ErrorResponse(errorMessage string, statusCode int) *Response {
	return &Response{
		Success:      false,
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
}

func SuccessResponse(payload interface{}) *Response {
	return &Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       payload,
	}
}
