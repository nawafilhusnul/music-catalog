package httpclient

import "net/http"

//go:generate mockgen -source=client.go -destination=client_mock.go -package=httpclient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	client HTTPClient
}

func NewClient(c HTTPClient) HTTPClient {
	return &client{client: c}
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
