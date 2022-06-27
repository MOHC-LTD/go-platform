package response

import (
	"go-platform/logger"
	"net/http"
)

// NotImplemented is thrown when a request method has not been implemented
type NotImplemented struct {
	Message string
}

func (err NotImplemented) Error() string {
	if err.Message == "" {
		return "method not implemented"
	}
	return err.Message
}

func (err NotImplemented) Response() any {
	return buildErrorResponse(err.Error())
}

func (err NotImplemented) Code() int {
	return http.StatusNotImplemented
}

func (err NotImplemented) Unwrap() error {
	return nil
}

func (err NotImplemented) Severity() logger.Severity {
	return logger.SeverityDebug
}
