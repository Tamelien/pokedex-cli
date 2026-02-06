package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(pageUrl *string) (LocationArea, error) {

	url := c.baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		var locationArea LocationArea
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, body)

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
