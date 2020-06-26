package translation

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

type Translator struct {
	Uni *ut.UniversalTranslator
}

func NewTranslator(locales ...locales.Translator) (*Translator, error) {
	fallback := en.New()

	trans := ut.New(fallback, locales...)
	if err := trans.Import(ut.FormatJSON, "translators"); err != nil {
		return nil, err
	}

	if err := trans.VerifyTranslations(); err != nil {
		return nil, err
	}

	return &Translator{trans}, nil
}
