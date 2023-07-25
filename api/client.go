package api

import (
	"net/http"
	"time"
)

// Client is in charge of interacting with a pihole server
// source : https://bitfieldconsulting.com/golang/api-client
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(url string, key string) *Client {
	return &Client{
		APIKey:  key,
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
