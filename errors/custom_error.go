package errors

import (
	"fmt"
	"net/http"
)

// AppError is a unified error structure for all services.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Common helpers
func New(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func BadRequest(msg string) *AppError {
	return New(http.StatusBadRequest, msg)
}

func Unauthorized(msg string) *AppError {
	return New(http.StatusUnauthorized, msg)
}

func Forbidden(msg string) *AppError {
	return New(http.StatusForbidden, msg)
}

func NotFound(msg string) *AppError {
	return New(http.StatusNotFound, msg)
}

func Conflict(msg string) *AppError {
	return New(http.StatusConflict, msg)
}

func Internal(msg string) *AppError {
	return New(http.StatusInternalServerError, msg)
}
