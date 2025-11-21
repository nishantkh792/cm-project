package http

import (
	"io"
	"net/http"
)

type HttpRequest struct{}

func NewHttp() *HttpRequest {
	return &HttpRequest{}
}

func (h *HttpRequest) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
