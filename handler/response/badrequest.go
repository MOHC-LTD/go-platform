package response

import (
	"go-platform/logger"
	"net/http"
)

// BadRequest is thrown when a request is formatted incorrectly
type BadRequest struct {
	Message string
	Err     error
}

func (err BadRequest) Error() string {
	if err.Message == "" {
		return "incorrect data payload"
	}
	return err.Message
}

func (err BadRequest) Response() any {
	return buildErrorResponse(err.Error())
}

func (err BadRequest) Code() int {
	return http.StatusBadRequest
}

func (err BadRequest) Unwrap() error {
	return err.Err
}

func (err BadRequest) Severity() logger.Severity {
	return logger.SeverityDebug
}
