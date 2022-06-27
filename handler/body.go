package handler

import (
	"encoding/json"
)

type responseBody struct {
	Error string `json:"error,omitempty"`
}

// marshal marshals any input into JSON-encoded bytes
func marshal(v any) ([]byte, error) {
	// Indenting can double body size!
	return json.Marshal(v)
}

// Bytes returns a byte representation of the response
func marshalResponse(v any) ([]byte, error) {
	jsonBytes, err := marshal(v)
	if err != nil {
		// Data or error has an unsupported value for JSON encoding, remove them from response
		jsonBytes, _ = marshal(responseBody{})
		return jsonBytes, err
	}

	return jsonBytes, nil
}
