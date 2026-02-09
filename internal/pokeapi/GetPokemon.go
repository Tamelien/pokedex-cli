package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {

	if pokemonName == "" {
		return Pokemon{}, fmt.Errorf("No pokemon to catch")
	}

	url := c.baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("%s is not a pokemon!", pokemonName)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
