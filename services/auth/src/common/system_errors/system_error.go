package system_errors

import (
	"net/http"
	"time"
)

type LogicalError struct {
	message string
	code    http.ConnState
}

type LogicalErrorResponse struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func (e LogicalError) Error() string {
	return e.message
}

func NewLogicalError(message string, code http.ConnState) *LogicalError {
	return &LogicalError{
		message: message,
		code:    code,
	}
}

func MapLogicalErrorToResponse(e *LogicalError) LogicalErrorResponse {
	return LogicalErrorResponse{
		Message: e.message,
		Date:    time.Now(),
	}
}
