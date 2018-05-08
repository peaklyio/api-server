package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	authorization Authorization
	configuration *ServerConfiguration
}

type ServerConfiguration struct {
	Address string
	Port    int
}

func NewClient(auth Authorization, cfg *ServerConfiguration) (*Client, error) {
	client := &Client{
		authorization: auth,
		configuration: cfg,
	}
	err := client.authorize()
	if err != nil {
		return nil, fmt.Errorf("unable to authorize client: %v", err)
	}
	return client, nil
}

func (c *Client) authorize() error {

	// Todo make this work lol

	return nil
}

func (c *Client) GET(endpoint string, params map[string]string) (*http.Response, error) {
	return nil, nil
}

func (c *Client) PUT(endpoint string, params map[string]string) (*http.Response, error) {
	return nil, nil
}

func (c *Client) PATCH(endpoint string, params map[string]string) (*http.Response, error) {
	return nil, nil
}

func (c *Client) DELETE(endpoint string, params map[string]string) (*http.Response, error) {
	return nil, nil
}
