package domain

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

var (
	Generic400 = &ErrorResponse{Code: http.StatusBadRequest, Message: "Invalid request parameters"}          // 400
	Generic401 = &ErrorResponse{Code: http.StatusUnauthorized, Message: "Unauthorized access"}             // 401
	Generic403 = &ErrorResponse{Code: http.StatusForbidden, Message: "Access forbidden"}                   // 403
	Generic404 = &ErrorResponse{Code: http.StatusNotFound, Message: "Resource not found"}                  // 404
	Generic422 = &ErrorResponse{Code: http.StatusUnprocessableEntity, Message: "Validation failed"}        // 422
	Generic429 = &ErrorResponse{Code: http.StatusTooManyRequests, Message: "Rate limit exceeded"}          // 429
)

var (
	ErrBadRequest      = Generic400
	ErrUnauthorized    = Generic401
	ErrForbidden       = Generic403
	ErrNotFound        = Generic404
	ErrUnprocessable   = Generic422
	ErrTooManyRequests = Generic429
)