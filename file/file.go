package file

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateTemporaryFile creates a new temporary file
func CreateTemporaryFile(filename string) (*os.File, error) {
	// Create temporary file to save download
	file, err := ioutil.TempFile(os.TempDir(), filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// DownloadFile downloads a file to the contents of a specified file, given a url
func DownloadFile(filename string, url string) (*os.File, error) {
	// Download file
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create temporary file to save download
	file, err := CreateTemporaryFile(filename)
	if err != nil {
		return nil, err
	}

	// Transfer to temporary file
	_, err = io.Copy(file, resp.Body)

	// Reset seek head
	file.Seek(0, io.SeekStart)

	return file, err
}
