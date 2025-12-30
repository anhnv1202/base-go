// Package apperror provides application-wide error types with proper HTTP status code mapping.
// These errors are used across all layers and provide consistent error handling.
package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

// ErrorCode represents application error codes
type ErrorCode string

const (
	// General errors
	ErrCodeInternal     ErrorCode = "INTERNAL_ERROR"
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeConflict     ErrorCode = "CONFLICT"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeBadRequest   ErrorCode = "BAD_REQUEST"

	// Domain-specific errors
	ErrCodeInvalidEmail    ErrorCode = "INVALID_EMAIL"
	ErrCodeInvalidPassword ErrorCode = "INVALID_PASSWORD"
	ErrCodeUserExists      ErrorCode = "USER_EXISTS"
	ErrCodeUserNotFound    ErrorCode = "USER_NOT_FOUND"
	ErrCodeInvalidToken    ErrorCode = "INVALID_TOKEN"
	ErrCodeTokenExpired    ErrorCode = "TOKEN_EXPIRED"
)

// AppError represents an application error with code, message, and HTTP status
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Details    string    `json:"details,omitempty"`
	HTTPStatus int       `json:"-"`
	Err        error     `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// WithDetails adds additional details to the error
func (e *AppError) WithDetails(details string) *AppError {
	e.Details = details
	return e
}

// WithError wraps an underlying error
func (e *AppError) WithError(err error) *AppError {
	e.Err = err
	return e
}

// New creates a new AppError
func New(code ErrorCode, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// Predefined errors for common cases

// ErrInternal returns an internal server error
func ErrInternal(message string) *AppError {
	return New(ErrCodeInternal, message, http.StatusInternalServerError)
}

// ErrValidation returns a validation error
func ErrValidation(message string) *AppError {
	return New(ErrCodeValidation, message, http.StatusBadRequest)
}

// ErrNotFound returns a not found error
func ErrNotFound(resource string) *AppError {
	return New(ErrCodeNotFound, fmt.Sprintf("%s not found", resource), http.StatusNotFound)
}

// ErrConflict returns a conflict error
func ErrConflict(message string) *AppError {
	return New(ErrCodeConflict, message, http.StatusConflict)
}

// ErrUnauthorized returns an unauthorized error
func ErrUnauthorized(message string) *AppError {
	return New(ErrCodeUnauthorized, message, http.StatusUnauthorized)
}

// ErrForbidden returns a forbidden error
func ErrForbidden(message string) *AppError {
	return New(ErrCodeForbidden, message, http.StatusForbidden)
}

// ErrBadRequest returns a bad request error
func ErrBadRequest(message string) *AppError {
	return New(ErrCodeBadRequest, message, http.StatusBadRequest)
}

// Domain-specific error constructors

// ErrInvalidEmail returns an invalid email error
func ErrInvalidEmail() *AppError {
	return New(ErrCodeInvalidEmail, "invalid email format", http.StatusBadRequest)
}

// ErrInvalidPassword returns an invalid password error
func ErrInvalidPassword(details string) *AppError {
	return New(ErrCodeInvalidPassword, "password does not meet requirements", http.StatusBadRequest).
		WithDetails(details)
}

// ErrUserExists returns a user exists error
func ErrUserExists() *AppError {
	return New(ErrCodeUserExists, "user already exists", http.StatusConflict)
}

// ErrUserNotFound returns a user not found error
func ErrUserNotFound() *AppError {
	return New(ErrCodeUserNotFound, "user not found", http.StatusNotFound)
}

// ErrInvalidToken returns an invalid token error
func ErrInvalidToken() *AppError {
	return New(ErrCodeInvalidToken, "invalid or malformed token", http.StatusUnauthorized)
}

// ErrTokenExpired returns a token expired error
func ErrTokenExpired() *AppError {
	return New(ErrCodeTokenExpired, "token has expired", http.StatusUnauthorized)
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr)
}

// AsAppError converts an error to AppError, or wraps it if not
func AsAppError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return ErrInternal("an unexpected error occurred").WithError(err)
}

// GetHTTPStatus returns the HTTP status code for an error
func GetHTTPStatus(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.HTTPStatus
	}
	return http.StatusInternalServerError
}
