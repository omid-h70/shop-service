package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type Success struct {
	statusCode int
	Response
}

func NewSuccess(msg string, result interface{}, status int) Success {
	return Success{
		statusCode: status,
		Response: Response{
			Message: msg,
			Result:  result,
		},
	}
}

func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	return json.NewEncoder(w).Encode(r.Response)
}
