package response

import (
	"github.com/MOHC-LTD/go-platform/logger"
	"net/http"
)

// Unauthorized is thrown when the authentication middleware fails to verify
type Unauthorized struct {
	Message string
	Err     error
}

func (err Unauthorized) Error() string {
	if err.Message == "" {
		return "unauthorized to access this resource"
	}
	return err.Message
}

func (err ServerError) Response() any {
	return buildErrorResponse(err.Error())
}

func (err Unauthorized) Code() int {
	return http.StatusUnauthorized
}

func (err Unauthorized) Unwrap() error {
	return err.Err
}

func (err Unauthorized) Severity() logger.Severity {
	return logger.SeverityDebug
}
