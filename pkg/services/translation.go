package services

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
)

var Uni *ut.UniversalTranslator

// initialize uni
func init() {
	fallback := en.New()

	Uni = ut.New(fallback, fallback)
	if err := Uni.Import(ut.FormatJSON, "translators"); err != nil {
		log.Fatal(errors.Wrap(err, "initialize translator"))
	}

	if err := Uni.VerifyTranslations(); err != nil {
		log.Fatal(errors.Wrap(err, "initialize translator"))
	}
}
