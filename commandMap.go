package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {

	locationArea, err := cfg.pokeapiClient.GetLocationList(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {

	url := cfg.prevLocationsURL
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationArea, err := cfg.pokeapiClient.GetLocationList(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, location := range locationArea.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationsURL = locationArea.Next
	cfg.prevLocationsURL = locationArea.Previous

	return nil
}
