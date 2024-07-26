package models

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) String() string {
	return fmt.Sprint("error: %s (code %d)", e.Message, e.Code)
}

const (
	ErrBadRequest          = http.StatusBadRequest
	ErrInternalServerError = http.StatusInternalServerError
)
