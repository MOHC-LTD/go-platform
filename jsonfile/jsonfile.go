package jsonfile

import (
	"encoding/json"
	"fmt"
	"github.com/MOHC-LTD/go-platform/jsonx"
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

// Read reads the named file and unmarshals the result into the value pointed to by T
func Read[T any](path string) (T, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return *new(T), err
	}

	return jsonx.Unmarshal[T](fileData)
}
