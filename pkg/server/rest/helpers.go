package rest

import (
	"demo/pkg/services"
	"encoding/json"
	"log"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"golang.org/x/text/language"
)

type errorStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func acceptLang(r *http.Request) ut.Translator {
	matcher := language.NewMatcher([]language.Tag{
		language.English,
		language.Vietnamese,
	})

	accept := r.Header.Get("Accept-Language")
	lang, _ := language.MatchStrings(matcher, accept)

	if trans, found := services.Uni.GetTranslator(lang.String()); found {
		return trans
	} else {
		return services.Uni.GetFallback()
	}
}

// response result
func response(status int, w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// response error
func responseError(status int, w http.ResponseWriter, r *http.Request, err error) {
	code := status
	message := http.StatusText(code)
	var details interface{}
	trans := acceptLang(r)

	if e, ok := err.(*services.ErrNotFound); ok {
		// not found error
		code = http.StatusNotFound
		message = e.Error()
	} else if e, ok := err.(*services.ErrValidate); ok {
		// validate error
		code = http.StatusBadRequest
		message = e.Error()
		details = e.Translate(trans)
	} else if e, ok := err.(*services.ErrInputInvalid); ok {
		// input invalid error
		code = http.StatusBadRequest
		message = e.Error()
		details = e.Translate(trans)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if code >= 500 {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(&errorStruct{
		Code:    code,
		Message: message,
		Details: details,
	})
}
