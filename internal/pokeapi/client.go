package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "github.com/akigithub888/pokedex/internal"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://pokeapi.co/api/v2",
		cache:      pokecache.NewCache(5 * time.Second),
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

	if data, ok := c.cache.Get(pageURL); ok {
		var cached LocationAreaList
		if err := json.Unmarshal(data, &cached); err == nil {
			return &cached, nil
		}
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

	c.cache.Add(pageURL, body)

	var data LocationAreaList
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
