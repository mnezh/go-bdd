package whatev

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Client is http.Client wrapper to ease test writing
type Client struct {
	httpClient http.Client
	baseURL    url.URL
}

// NewClient is client factory method
func NewClient(baseURL string) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &Client{http.Client{}, *parsedURL}, nil
}

// GetJSON send HTTP GET to uri relative to baseURL and parses output
func (c *Client) GetJSON(uri string) (*map[string]interface{}, *http.Response, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("GET", c.baseURL.ResolveReference(u).String(), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept-type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var parsedJSON map[string]interface{}
	json.NewDecoder(res.Body).Decode(&parsedJSON)
	return &parsedJSON, res, nil
}
