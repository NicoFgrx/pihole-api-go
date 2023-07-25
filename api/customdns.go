package api

import (
	"fmt"
	"net/http"
)

type DNSRecordParams struct {
	Domain string
	IP     string
}

type GetCustomDNSResponse struct {
	Data [][]string `json:"data"`
}

type PostCustomDNSResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (client *Client) GetCustomDNS() (*http.Response, error) {

	NewURL := fmt.Sprintf("%s/admin/api.php?customdns&action=get&auth=%s", client.BaseURL, client.APIKey)
	resp, err := client.HTTPClient.Get(NewURL)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (client *Client) AddCustomDNS(params *DNSRecordParams) (*http.Response, error) {

	NewURL := fmt.Sprintf("%s/admin/api.php?customdns&action=add&auth=%s&ip=%s&domain=%s", client.BaseURL, client.APIKey, params.IP, params.Domain)
	resp, err := client.HTTPClient.Get(NewURL)

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (client *Client) DeleteCustomDNS(params *DNSRecordParams) (*http.Response, error) {

	NewURL := fmt.Sprintf("%s/admin/api.php?customdns&action=delete&auth=%s&ip=%s&domain=%s", client.BaseURL, client.APIKey, params.IP, params.Domain)
	resp, err := client.HTTPClient.Get(NewURL)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
