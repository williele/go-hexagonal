package services

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

var Uni *ut.UniversalTranslator

// initialize uni
func init() {
	fallback := en.New()

	Uni = ut.New(fallback)
	if err := Uni.Import(ut.FormatJSON, "translators"); err != nil {
		log.Fatal(err)
	}

	if err := Uni.VerifyTranslations(); err != nil {
		log.Fatal(err)
	}
}
