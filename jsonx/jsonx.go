package jsonx

import "encoding/json"

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by T
func Unmarshal[T any](data []byte) (T, error) {
	var result T

	err := json.Unmarshal(data, &result)
	if err != nil {
		return *new(T), err
	}

	return result, nil
}
