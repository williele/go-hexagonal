package services

import (
	"log"
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// initialize validate
func init() {
	Validate = validator.New()

	// registe tag name as json field
	Validate.RegisterTagNameFunc(func(r reflect.StructField) string {
		name := strings.SplitN(r.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	// register translations
	registerTranslation()
}

// registe custom validators
func registerValidator() {
	Validate.RegisterValidation("price", func(fl validator.FieldLevel) bool {
		return fl.Field().Float() > 0
	})
}

// registe validator translation
func registerTranslation() {
	translations := []struct {
		tag       string
		transFunc validator.TranslationFunc
	}{
		// requried
		{tag: "required"},
		{tag: "price"},
	}

	// sign
	fallback := Uni.GetFallback()
	for _, t := range translations {
		if t.transFunc != nil {
			if err := Validate.RegisterTranslation(t.tag, fallback, registerFunc, t.transFunc); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := Validate.RegisterTranslation(t.tag, fallback, registerFunc, translateFunc); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func registerFunc(ut ut.Translator) error {
	return nil
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	field, err := ut.T(fe.Field())
	if err != nil {
		field = fe.Field()
	}

	t, err := ut.T(fe.Tag(), field)
	if err != nil {
		return fe.(error).Error()
	}
	return t
}
