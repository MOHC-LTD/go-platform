package handler

import (
	"fmt"
	"go-platform/handler/response"
	log "go-platform/logger"
	"net/http"
)

type routerError interface {
	Error() string
	Unwrap() error
	Severity() log.Severity
	Response() any
	Code() int
}

func respond(logger log.Logger, w http.ResponseWriter, r *http.Request, res any, statusCode int) {
	jsonBytes, err := marshalResponse(res)
	if err != nil {
		// jsonBytes will still have JSON encoded bytes but without data/error fields
		Log(logger, r, log.SeverityError, fmt.Errorf("marshalling json response: %w", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}

func successResponse(logger log.Logger, w http.ResponseWriter, r *http.Request, result any) {
	respond(logger, w, r, result, http.StatusOK)
}

func errorResponse(logger log.Logger, w http.ResponseWriter, r *http.Request, err error) {
	// Default to internal server error
	routerRes, ok := err.(routerError)
	if !ok {
		routerRes = response.ServerError{Err: err}
	}

	// Log error with path if there was one to unwrap
	if internalErr := routerRes.Unwrap(); internalErr != nil {
		Log(logger, r, routerRes.Severity(), internalErr)
	}

	respond(logger, w, r, routerRes.Response(), routerRes.Code())
}
