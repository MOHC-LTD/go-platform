package response

import (
	"lod-product/internal/platform/logger"
	"net/http"
)

// NotFound is thrown when a request is not found
type NotFound struct {
	Message string
}

func (err NotFound) Error() string {
	if err.Message == "" {
		return "resource not found"
	}
	return err.Message
}

func (err NotFound) Response() any {
	return buildErrorResponse(err.Error())
}

func (err NotFound) Code() int {
	return http.StatusNotFound
}

func (err NotFound) Unwrap() error {
	return nil
}

func (err NotFound) Severity() logger.Severity {
	return logger.SeverityDebug
}
