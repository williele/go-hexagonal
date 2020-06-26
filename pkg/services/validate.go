package services

import (
	"demo/pkg/translation"
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validate *validator.Validate
}

func NewValidator(translator *translation.Translator) *Validator {
	validate := validator.New()

	// registe tag name as json field
	validate.RegisterTagNameFunc(func(r reflect.StructField) string {
		name := strings.SplitN(r.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	// register translations
	registerTranslation(validate, translator.Uni)

	return &Validator{
		Validate: validate,
	}
}

func registerTranslation(validate *validator.Validate, uni *ut.UniversalTranslator) error {
	translations := []struct {
		tag       string
		transFunc validator.TranslationFunc
	}{
		// requried
		{
			tag: "required",
			transFunc: func(ut ut.Translator, fe validator.FieldError) string {
				message, _ := ut.T(fe.Tag(), fe.Field())
				return message
			},
		},
	}

	fallback := uni.GetFallback()
	for _, t := range translations {
		if err := validate.RegisterTranslation(t.tag, fallback, func(ut.Translator) error {
			return nil
		}, t.transFunc); err != nil {
			return err
		}
	}

	return nil
}
