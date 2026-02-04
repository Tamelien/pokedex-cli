package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(url string) (LocationArea, error) {

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

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
