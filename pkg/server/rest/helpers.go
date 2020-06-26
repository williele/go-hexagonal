package rest

import (
	"encoding/json"
	"net/http"
)

type errorStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func response(status int, w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func responseError(status int, w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&errorStruct{
		Code:    status,
		Message: err.Error(),
	})
}
