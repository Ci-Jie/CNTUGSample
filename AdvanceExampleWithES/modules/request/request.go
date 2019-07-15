package request

import (
	"io"
	"net/http"
)

// Request is responsible to send a request to specific URL.
func Request(method string, url string, data io.Reader) error {
	req, err := http.NewRequest(method, url, data)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
