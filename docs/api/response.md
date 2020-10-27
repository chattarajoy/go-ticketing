# API Contracts and Responses

## Standard Response Format

```go
type Response struct {
	Success      bool        `json:"success"`
	StatusCode   int         `json:"status_code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}
```