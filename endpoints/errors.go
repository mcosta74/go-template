package endpoints

import (
	"fmt"
	"net/http"
)

type ApiError struct {
	code int
	err  error
}

func (e ApiError) StatusCode() int {
	return e.code
}

func (e ApiError) Error() string {
	return e.err.Error()
}

func (e ApiError) Unwrap() error {
	return e.err
}

func NewApiError(code int, err error) *ApiError {
	return &ApiError{code: code, err: err}
}

func NewDecodeError(err error) *ApiError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("decode error: %w", err))
}

func NewNotFoundError(err error) *ApiError {
	return NewApiError(http.StatusNotFound, err)
}

func NewConflictError(err error) *ApiError {
	return NewApiError(http.StatusConflict, err)
}
