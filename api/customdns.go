package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

func (client *Client) GetCustomDNS() ([]DNSRecordParams, error) {

	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return nil, err
	}

	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "get")

	NewURL.RawQuery = v.Encode()

	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// format output to return array of DNSRecord

	var post GetCustomDNSResponse

	var customdns_lst []DNSRecordParams

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured : %s", err)
	}

	for i := 0; i < len(post.Data); i++ {
		customdns_lst = append(customdns_lst, DNSRecordParams{
			Domain: post.Data[i][0],
			IP:     post.Data[i][1],
		})
	}

	return customdns_lst, nil

}

func (client *Client) AddCustomDNS(params *DNSRecordParams) (*http.Response, error) {

	// NewURL := fmt.Sprintf("%s/admin/api.php?customdns&action=add&auth=%s&ip=%s&domain=%s", client.BaseURL, client.APIKey, params.IP, params.Domain)
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return nil, err
	}

	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "add")
	v.Set("ip", params.IP)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	resp, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (client *Client) DeleteCustomDNS(params *DNSRecordParams) (*http.Response, error) {

	// NewURL := fmt.Sprintf("%s/admin/api.php?customdns&action=delete&auth=%s&ip=%s&domain=%s", client.BaseURL, client.APIKey, params.IP, params.Domain)

	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return nil, err
	}

	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "delete")
	v.Set("ip", params.IP)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	resp, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return nil, err
	}

	return resp, nil

}
