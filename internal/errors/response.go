package errors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error     *AppError `json:"error"`
	RequestID string    `json:"request_id,omitempty"`
	Timestamp string    `json:"timestamp,omitempty"`
}

func WriteErrorResponse(w http.ResponseWriter, err error, requestID string) {
	var appErr *AppError

	if IsAppError(err) {
		appErr, _ = AsAppError(err)
	} else {
		appErr = NewInternalError("UNKNOWN_ERROR", "An unexpected error occured", err)
	}

	response := ErrorResponse{
		Error:     appErr,
		RequestID: requestID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.HTTPStatus())
	json.NewEncoder(w).Encode(response)
}

func WriteValidationError(w http.ResponseWriter, code, message string, details map[string]interface{}, requestID string) {
	appErr := NewValidationError(code, message, details)
	WriteErrorResponse(w, appErr, requestID)
}

func WriteNotFoundError(w http.ResponseWriter, code, message string, requestID string) {
	appErr := NewNotFoundError(code, message)
	WriteErrorResponse(w, appErr, requestID)
}
