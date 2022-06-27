package handler

import (
	"encoding/json"
	"github.com/MOHC-LTD/go-platform/handler/response"
	log "github.com/MOHC-LTD/go-platform/logger"
	"io"
	"net/http"
)

// JSONFunc is a handler function called by the JSON Handler, with payload being unmarshalled prior to call
type JSONFunc[T any] func(*http.Request, T) (any, error)

// NewJSON returns a handler func that handles JSON request payloads
func NewJSON[T any](logger log.Logger, handlerFunc JSONFunc[T]) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var payload *T

		// Unmarshal contents of body
		err := json.NewDecoder(r.Body).Decode(&payload)
		switch {
		case err == io.EOF:
			// Allow empty request bodies
			payload = new(T)
		case err != nil:
			errorResponse(logger, rw, r, response.BadRequest{Err: err})
			return
		}

		// Call handler function
		result, err := handlerFunc(r, *payload)
		if err != nil {
			errorResponse(logger, rw, r, err)
			return
		}

		successResponse(logger, rw, r, result)
	}
}
