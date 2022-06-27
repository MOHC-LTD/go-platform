package handler

import (
	"github.com/MOHC-LTD/go-platform/logger"
	"net/http"
)

type Middleware func(next http.Handler) http.Handler
type MiddlewareFunc func(*http.Request) (*http.Request, error)

// NewMiddleware returns a handler func that handles requests
func NewMiddleware(logger logger.Logger, handlerFunc MiddlewareFunc) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			// Call handler function
			r, err := handlerFunc(r)
			if err != nil {
				errorResponse(logger, rw, r, err)
				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}
