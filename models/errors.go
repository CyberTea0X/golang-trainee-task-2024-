package models

import "errors"

var (
	ErrInvalidJson  = errors.New("invalid json")
	ErrInvalidQuery = errors.New("invalid query")
	ErrInternal     = errors.New("internal server error")
	ErrInvalidId    = errors.New("invalid banner id in url path")
)

var (
	ErrSqlNoRowsDeleted = errors.New("no rows were deleted")
)
