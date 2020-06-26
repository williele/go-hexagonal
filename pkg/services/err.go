package services

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// not found error
type ErrNotFound struct {
	Name string
}

func NewErrNotFound(name string) error {
	return &ErrNotFound{name}
}

func (e *ErrNotFound) Error() string {
	return e.Name + " not found"
}

// validate error
type ErrValidate struct {
	Input   string
	Details validator.ValidationErrors
}

func NewErrValidate(input string, details validator.ValidationErrors) error {
	return &ErrValidate{input, details}
}

func (e *ErrValidate) Error() string {
	return e.Input + " input invalid"
}

func (e *ErrValidate) Translate(trans ut.Translator) map[string]interface{} {
	result := make(map[string]interface{})
	for _, field := range e.Details {
		result[field.Field()] = field.Translate(trans)
	}

	return result
}

// input invalid error
type ErrInputInvalid struct {
	Input   string
	Details map[string]string
}

func NewErrInputInvalid(input string, details map[string]string) error {
	return &ErrInputInvalid{input, details}
}

func (e *ErrInputInvalid) Error() string {
	return e.Input + " input invalid"
}

func (e *ErrInputInvalid) Translate(trans ut.Translator) map[string]interface{} {
	result := make(map[string]interface{})
	for field, key := range e.Details {
		name, err := trans.T(field)
		if err != nil {
			name = field
		}
		result[field], _ = trans.T(key, name)
	}

	return result
}
