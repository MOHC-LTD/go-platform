package handler

import (
	log "lod-product/internal/platform/logger"
	"net/http"
)

type Func func(*http.Request) (any, error)
type WriterFunc func(rw http.ResponseWriter, r *http.Request) error

// New returns a handler func that handles requests
func New(logger log.Logger, handlerFunc Func) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Call handler function
		result, err := handlerFunc(r)
		if err != nil {
			errorResponse(logger, rw, r, err)
			return
		}

		successResponse(logger, rw, r, result)
	}
}

// NewWriter returns a handler func that handles requests and allows the handler to write the response
func NewWriter(logger log.Logger, handlerFunc WriterFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Call handler function
		err := handlerFunc(rw, r)
		if err != nil {
			errorResponse(logger, rw, r, err)
			return
		}
	}
}
