package server

import (
	"errors"
	"net/http"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidIdValue  = errors.New("invalid ID value")
)

func ErrorMapper(err error) (int, error) {
	switch err {
	case ErrProductNotFound:
		return http.StatusNotFound, err
	case ErrInvalidIdValue:
		return http.StatusBadRequest, err
	default:
		return http.StatusInternalServerError, err
	}
}
