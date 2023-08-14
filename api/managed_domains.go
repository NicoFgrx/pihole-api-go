package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type GetManagedDomainsResponse struct {
	Data []ManagedDomain `json:"data"`
}

type PostManagedDomainResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ManagedDomain struct {
	Id            int
	Type          int
	Domain        string
	Enabled       int
	Date_added    int
	Date_modified int
	Comment       string
	Groups        []int
}

var Categories = map[int]string{
	0: "white",
	1: "black",
	2: "regex_white",
	3: "regex_black",
}

func IsValidCategory(category string) bool {
	switch category {
	case
		"white",
		"black",
		"regex_white",
		"regex_black":
		return true
	}
	return false
}

func (client *Client) GetManagedDomainsFromCategory(category string) ([]ManagedDomain, error) {

	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return nil, err
	}

	// check user input
	if !IsValidCategory(category) {
		return nil, fmt.Errorf("Invalid category founded : %s", category)
	}
	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)
	v.Set("list", category)

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// format output to return slice of ManagedDomain

	var get GetManagedDomainsResponse
	var managed_domain_lst []ManagedDomain

	// Decode the json slice
	if err := json.NewDecoder(res.Body).Decode(&get); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// Build slice of ManagedDomain
	for i := 0; i < len(get.Data); i++ {
		managed_domain_lst = append(managed_domain_lst, get.Data[i])
	}

	// Return the slice
	return managed_domain_lst, nil

}

func (client *Client) GetAllManagedDomains() ([]ManagedDomain, error) {

	var all_managed_domain_lst []ManagedDomain

	for _, v := range Categories {
		curr_managed_domain_lst, err := client.GetManagedDomainsFromCategory(v)

		if err != nil {
			return nil, err
		}
		all_managed_domain_lst = append(all_managed_domain_lst, curr_managed_domain_lst...)

	}

	return all_managed_domain_lst, nil

}

func (client *Client) AddManagedDomains(category, domain string) error {
	// Build the url
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// check user input
	if !IsValidCategory(category) {
		return fmt.Errorf("Invalid category founded : %s", category)
	}

	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)
	v.Set("list", category)
	v.Set("add", domain)

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var status PostManagedDomainResponse

	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)
	}

	return nil

}

func (client *Client) SubManagedDomains(category, domain string) error {
	// Build the url
	NewURL, err := url.Parse(client.BaseURL)
	if err != nil {
		return err
	}

	// check user input
	if !IsValidCategory(category) {
		return fmt.Errorf("Invalid category founded : %s", category)
	}

	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)
	v.Set("list", category)
	v.Set("sub", domain)

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var status PostManagedDomainResponse

	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// If the API inform us that the resource cannot be created, return an error
	if !status.Success {
		return fmt.Errorf(status.Message)
	}

	return nil

}
