package main

import "github.com/Tamelien/pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL string
	prevLocationsURL string
}

func main() {
	client := pokeapi.NewClient()
	cfg := &config{
		pokeapiClient:    client,
		nextLocationsURL: "https://pokeapi.co/api/v2/location-area/",
		prevLocationsURL: "",
	}
	repl(cfg)
}
