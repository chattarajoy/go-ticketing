package json

import (
	"encoding/json"
	"net/http"
)

func WriteResult(v interface{}, writer http.ResponseWriter, statusCode int) {
	jData, _ := json.Marshal(v)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_, _ = writer.Write(jData)
}
