package jsonfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Unmarshal parses the JSON-encoded contents of a file into a value generically typed
func Unmarshal[T any](filePath string) (T, error) {
	pointer := new(T)

	f, err := os.Open(filePath)
	if err != nil {
		return *pointer, fmt.Errorf("opening config file: %w", err)
	}
	defer f.Close()

	jsonBytes, err := io.ReadAll(f)
	if err != nil {
		return *pointer, fmt.Errorf("reading config file: %w", err)
	}

	err = json.Unmarshal(jsonBytes, &pointer)
	if err != nil {
		return *pointer, fmt.Errorf("unmarshalling config file: %w", err)
	}

	return *pointer, nil
}
