package models

import "errors"

var (
	ErrInvalidJson = errors.New("invalid json")
	ErrInternal    = errors.New("internal server error")
	ErrInvalidId   = errors.New("invalid banner id in url path")
)
