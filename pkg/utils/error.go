package utils

import "errors"

var (
	ErrProductExists        = errors.New("product name already exists")
	ErrInvalidName          = errors.New("product name is invalid or empty")
	ErrInvalidPrice         = errors.New("price must be zero or more")
	ErrInvalidQuantity      = errors.New("quantity must be zero or more")
	ErrInvalidSortField     = errors.New("invalid sort key")
	ErrInvalidSortDirection = errors.New("invalid sort direction")
	ErrDescriptionTooLong   = errors.New("description too long (max 500)")
	ErrNameAlphaNumeric     = errors.New("product name must be alphanumeric and readable")
)
