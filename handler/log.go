package handler

import (
	"fmt"
	log "lod-product/internal/platform/logger"
	"net/http"
)

// Log logs an error encountered during the processing of a request
func Log(logger log.Logger, r *http.Request, severity log.Severity, err error) {
	errMessage := fmt.Sprintf("%v %v: %v", r.Method, r.URL.Path, err)
	switch severity {
	case log.SeverityDebug:
		logger.Debug(errMessage)
	case log.SeverityInfo:
		logger.Info(errMessage)
	case log.SeverityError:
		logger.Error(errMessage)
	}
}
