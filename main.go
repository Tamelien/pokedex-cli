package main

import (
	"time"

	"github.com/Tamelien/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	client := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: client,
	}
	repl(cfg)
}
