package errors

import (
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "VALIDATION"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeConflict     ErrorType = "CONFLICT"
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	ErrorTypeForbidden    ErrorType = "FORBIDDEN"
	ErrorTypeInternal     ErrorType = "INTERNAL"
	ErrorTypeBadRequest   ErrorType = "BAD_REQUEST"
)

type AppError struct {
	Type    ErrorType              `json:"type"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
	Cause   error                  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Code, e.Message, e.Cause)
	}

	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) HTTPStatus() int {
	switch e.Type {
	case ErrorTypeValidation, ErrorTypeBadRequest:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeConflict:
		return http.StatusConflict
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

func AsAppError(err error) (*AppError, bool) {
	if appErr, ok := err.(*AppError); ok {
		return appErr, true
	}
	return nil, false
}

func NewValidationError(code, message string, details map[string]interface{}) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Code:    code,
		Message: message,
		Details: details,
	}
}

func NewNotFoundError(code, message string) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Code:    code,
		Message: message,
	}
}

func NewConflictError(code, message string) *AppError {
	return &AppError{
		Type:    ErrorTypeConflict,
		Code:    code,
		Message: message,
	}
}

func NewUnauthorizedError(code, message string) *AppError {
	return &AppError{
		Type:    ErrorTypeUnauthorized,
		Code:    code,
		Message: message,
	}
}

func NewForbiddenError(code, message string) *AppError {
	return &AppError{
		Type:    ErrorTypeForbidden,
		Code:    code,
		Message: message,
	}
}

func NewInternalError(code, message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeInternal,
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func NewBadRequestError(code, message string) *AppError {
	return &AppError{
		Type:    ErrorTypeBadRequest,
		Code:    code,
		Message: message,
	}
}

func WrapError(err error, code, message string) *AppError {
	if appErr, ok := AsAppError(err); ok {
		appErr.Message = message
		return appErr
	}

	return &AppError{
		Type:    ErrorTypeInternal,
		Code:    code,
		Message: message,
		Cause:   err,
	}
}
