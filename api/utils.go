package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func (client *Client) GetStatus() (string, error) {
	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return "", err
	}

	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)
	v.Set("status", "")

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// format output to return slice of CNAMERecord

	var post StatusResponse

	// Decode the slice of slice
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// Return the slice
	return post.Status, nil

}

func (client *Client) DisableBlocking(n ...int64) (string, error) {
	var time int64

	if len(n) != 0 && len(n) != 1 {
		return "", fmt.Errorf("Invalid parameter founed : %d, need 0 or 1 argument", len(n))
	}

	if len(n) == 1 {
		time = n[0]
	}

	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return "", err
	}

	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)

	v.Set("disable", "")

	if n != nil {
		v.Set("disable", strconv.FormatInt(time, 10))

	}

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	var post StatusResponse

	// Decode the slice of slice
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	return post.Status, nil

}

func (client *Client) EnableBlocking() (string, error) {

	// Build the url
	NewURL, err := url.Parse(client.BaseURL)

	if err != nil {
		return "", err
	}

	// Set the url values
	v := url.Values{}

	v.Set("auth", client.APIKey)
	v.Set("enable", "")

	NewURL.RawQuery = v.Encode()

	// Send GET request
	res, err := client.HTTPClient.Get(NewURL.String())

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// format output to return slice of CNAMERecord

	var post StatusResponse

	// Decode the slice of slice
	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		log.Fatalf("An error occured while decode the JSON : %s", err)
	}

	// Return the slice
	return post.Status, nil

}
