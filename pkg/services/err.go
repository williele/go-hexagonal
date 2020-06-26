package services

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInputInvalid = errors.New("input invalid")
)
