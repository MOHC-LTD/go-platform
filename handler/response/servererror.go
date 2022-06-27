package response

import (
	"lod-product/internal/platform/logger"
	"net/http"
)

// ServerError is thrown when the server encounters an error during request processing
type ServerError struct {
	Message string
	Err     error
}

func (err ServerError) Error() string {
	if err.Message == "" {
		return "internal server error"
	}
	return err.Message
}

func (err Unauthorized) Response() any {
	return buildErrorResponse(err.Error())
}

func (err ServerError) Code() int {
	return http.StatusInternalServerError
}

func (err ServerError) Unwrap() error {
	return err.Err
}

func (err ServerError) Severity() logger.Severity {
	return logger.SeverityError
}
