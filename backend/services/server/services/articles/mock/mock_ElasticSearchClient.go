package repository

import (
	"io"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type MockTransport struct {
	Response    *http.Response
	RoundTripFn func(req *http.Request) (*http.Response, error)
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFn(req)
}

func NewMockElasticClient() (*elasticsearch.Client, error) {
	mocktrans := MockTransport{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{}`)),
		},
	}
	mocktrans.RoundTripFn = func(req *http.Request) (*http.Response, error) { return mocktrans.Response, nil }

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Transport: &mocktrans,
	})
	if err != nil {
		return client, err
	}

	return client, nil
}