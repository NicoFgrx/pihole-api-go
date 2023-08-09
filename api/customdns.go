package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

// Ask the pihole API for all existing dns records
// Return a slice of DNSRecordParams
// If an error has occured, earlier or with the GET request, the error is return
func (client *Client) GetAllCustomDNS() ([]DNSRecordParams, error) {

	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return nil, err
	}

	// Set the url values
	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "get")

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// format output to return slice of DNSRecord

	var post GetCustomDNSResponse
	var customdns_lst []DNSRecordParams

	// Decode the slice of slice
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// Build slice of DNSRecordParams
	for i := 0; i < len(post.Data); i++ {
		customdns_lst = append(customdns_lst, DNSRecordParams{
			Domain: post.Data[i][0],
			IP:     post.Data[i][1],
		})
	}

	// Return the slice
	return customdns_lst, nil

}

// Ask the pihole API for all existing dns records then search if the given domain exist
// Return a DNSRecordParams if it's find in the slice
// If an error has occured, earlier or if the domain is not founded, the error is return
func (client *Client) GetCustomDNS(domain string) (DNSRecordParams, error) {

	// Fetch all dns records
	customdns_lst, err := client.GetAllCustomDNS()
	if err != nil {
		return DNSRecordParams{}, err
	}

	// Return item if the domain exist in the slice
	for _, item := range customdns_lst {
		if item.Domain == domain {
			return item, nil
		}
	}

	// Return an error if the domain doesn't exist in the slice
	return DNSRecordParams{}, errors.New("Records not found")

}

// Ask the pihole API to create a new DNS Record
// Return true if the operation is a success, return an error if not
func (client *Client) AddCustomDNS(params *DNSRecordParams) error {

	// Bulid the URL
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// Set the url values
	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "add")
	v.Set("ip", params.IP)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	// Send GET request
	resp, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return err
	}

	// Check if API Response is true or false
	var status PostCustomDNSResponse

	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)

	}

	return nil

}

// Ask the pihole API to delete a new DNS Record
// Return true if the operation is a success, return an error if not
func (client *Client) DeleteCustomDNS(params *DNSRecordParams) error {

	// Build the URL
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// Set url values
	v := url.Values{}

	v.Set("customdns", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "delete")
	v.Set("ip", params.IP)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	// Set GET request
	resp, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return err
	}

	// Check if API Response is true or false
	var status PostCustomDNSResponse

	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)
	}

	return nil

}

// Ask the pihole API to delete, then re-create the dns record
// If an error has occured, the error is returned
func (client *Client) UpdateCustomDNS(domain string, params *DNSRecordParams) error {

	// Get current record to delete
	data, err := client.GetCustomDNS(domain)
	if err != nil {
		return err
	}

	// Delete current record
	err = client.DeleteCustomDNS(&data)
	if err != nil {
		return err
	}

	// Create new record
	err = client.AddCustomDNS(params)
	if err != nil {
		return err
	}

	return nil

}
