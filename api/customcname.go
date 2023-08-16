package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type CNAMERecordParams struct {
	Domain string
	Target string
}

type GetCustomCNAMEResponse struct {
	Data [][]string `json:"data"`
}

type PostCustomCNAMEResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Ask the pihole API for all existing CNAME records
// Return a slice of CNAMERecordParams
// If an error has occured, earlier or with the GET request, the error is return
func (client *Client) GetAllCustomCNAME() ([]CNAMERecordParams, error) {

	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return nil, err
	}

	// Set the url values
	v := url.Values{}

	v.Set("customcname", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "get")

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// format output to return slice of CNAMERecord

	var post GetCustomCNAMEResponse
	var customCNAME_lst []CNAMERecordParams

	// Decode the slice of slice
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// Build slice of CNAMERecordParams
	for i := 0; i < len(post.Data); i++ {
		customCNAME_lst = append(customCNAME_lst, CNAMERecordParams{
			Domain: post.Data[i][0],
			Target: post.Data[i][1],
		})
	}

	// Return the slice
	return customCNAME_lst, nil

}

// Ask the pihole API for all existing CNAME records then search if the given domain exist
// Return a CNAMERecordParams if it's find in the slice
// If an error has occured, earlier or if the domain is not founded, the error is return
func (client *Client) GetCustomCNAME(domain string) (CNAMERecordParams, error) {

	// Fetch all CNAME records
	customCNAME_lst, err := client.GetAllCustomCNAME()
	if err != nil {
		return CNAMERecordParams{}, err
	}

	// Return item if the domain exist in the slice
	for _, item := range customCNAME_lst {
		if item.Domain == domain {
			return item, nil
		}
	}

	// Return an error if the domain doesn't exist in the slice
	return CNAMERecordParams{}, fmt.Errorf("CNAME %s not found", domain)

}

// Ask the pihole API to create a new CNAME Record
// Return true if the operation is a success, return an error if not
func (client *Client) AddCustomCNAME(params *CNAMERecordParams) error {

	// Bulid the URL
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// Set the url values
	v := url.Values{}

	v.Set("customcname", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "add")
	v.Set("target", params.Target)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Check if API Response is true or false
	var status PostCustomCNAMEResponse

	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)

	}

	return nil

}

// Ask the pihole API to delete a new CNAME Record
// Return true if the operation is a success, return an error if not
func (client *Client) DeleteCustomCNAME(params *CNAMERecordParams) error {

	// Build the URL
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// Set url values
	v := url.Values{}

	v.Set("customcname", "")
	v.Set("auth", client.APIKey)
	v.Set("action", "delete")
	v.Set("target", params.Target)
	v.Set("domain", params.Domain)

	NewURL.RawQuery = v.Encode()

	// Set GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Check if API Response is true or false
	var status PostCustomCNAMEResponse

	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)
	}

	return nil

}

// Ask the pihole API to delete, then re-create the CNAME record
// If an error has occured, the error is returned
func (client *Client) UpdateCustomCNAME(domain string, params *CNAMERecordParams) error {

	// Get current record to delete
	data, err := client.GetCustomCNAME(domain)
	if err != nil {
		return err
	}

	// Delete current record
	err = client.DeleteCustomCNAME(&data)
	if err != nil {
		return err
	}

	// Create new record
	err = client.AddCustomCNAME(params)
	if err != nil {
		return err
	}

	return nil

}
