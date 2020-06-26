package rest

import (
	"demo/pkg/services"
	"encoding/json"
	"net/http"
)

type errorStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// response result
func response(status int, w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// response error
func responseError(status int, w http.ResponseWriter, r *http.Request, err error) {
	code := status
	message := http.StatusText(code)
	var details interface{}

	if e, ok := err.(*services.ErrNotFound); ok {
		// not found error
		code = http.StatusNotFound
		message = e.Error()
	} else if e, ok := err.(*services.ErrValidate); ok {
		// validate error
		code = http.StatusBadRequest
		message = e.Error()
		details = e.Translate(services.Uni.GetFallback())
	} else if e, ok := err.(*services.ErrInputInvalid); ok {
		// input invalid error
		code = http.StatusBadRequest
		message = e.Error()
		details = e.Translate(services.Uni.GetFallback())
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&errorStruct{
		Code:    code,
		Message: message,
		Details: details,
	})
}
