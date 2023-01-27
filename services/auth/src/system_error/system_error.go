package system_error

import "net/http"

type logicalError struct {
	message string
	code    http.ConnState
}

type LogicalError *logicalError

func (e *logicalError) Error() string {
	return e.message
}

func NewLogicalError(message string, code http.ConnState) LogicalError {
	return &logicalError{
		message: message,
		code:    code,
	}
}
