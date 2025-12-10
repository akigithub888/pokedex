package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://pokeapi.co/api/v2",
	}
}

type LocationAreaList struct {
	Count    int                      `json:"count"`
	Next     *string                  `json:"next"`
	Previous *string                  `json:"previous"`
	Results  []LocationAreaListResult `json:"results"`
}

type LocationAreaListResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c *Client) ListLocationAreas(pageURL string) (*LocationAreaList, error) {
	if pageURL == "" {
		pageURL = c.baseURL + "/location-area"
	}
	res, err := c.httpClient.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data LocationAreaList
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
