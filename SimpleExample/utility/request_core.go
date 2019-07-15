package utility

import (
	"fmt"
	"io"

	"net/http"
	"strings"
)

// SendRequest ...
func SendRequest(r *http.Request) (*http.Response, error) {

	body := parseBody(r)

	request, err := http.NewRequest(r.Method, fmt.Sprintf("http://172.17.1.100:7480%s", r.URL), body)

	if err != nil {
		return nil, err
	}

	if r.Header != nil {
		for index, value := range r.Header {
			request.Header.Add(index, strings.Join(value, " "))
		}
	}

	request.Host = r.Host

	client := &http.Client{}

	return client.Do(request)
}

func parseBody(r *http.Request) io.Reader {

	length := r.ContentLength

	if length == 0 {
		return strings.NewReader("")
	}

	body := make([]byte, length)
	r.Body.Read(body)

	return strings.NewReader(string(body))
}
