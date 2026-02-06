package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationEncounters(location string) (LocationEncounters, error) {

	if location == "" {
		return LocationEncounters{}, fmt.Errorf("No location provided")
	}

	url := c.baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		var locationEncounters LocationEncounters
		if err := json.Unmarshal(val, &locationEncounters); err != nil {
			return LocationEncounters{}, err
		}
		return locationEncounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationEncounters{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationEncounters{}, err
	}
	c.cache.Add(url, body)

	var locationEncounters LocationEncounters
	if err := json.Unmarshal(body, &locationEncounters); err != nil {
		return LocationEncounters{}, err
	}

	return locationEncounters, nil
}
