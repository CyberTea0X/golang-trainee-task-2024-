package models

import "errors"

var (
	ErrInvalidJson    = errors.New("invalid json")
	ErrInternalServer = errors.New("internal server error")
)
