package errors

import "errors"

var (
	ErrInvalidLine   = errors.New("invalid line")
	ErrInvalidFloat  = errors.New("invalid float")
	ErrUnknownSensor = errors.New("unknown sensor")
)
