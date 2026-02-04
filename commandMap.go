package main

import (
	"fmt"
)

func commandMap(cfg *config) error {

	url := cfg.nextLocationsURL
	if url == "" {
		fmt.Println("No location pages available.")
		return nil
	}

	locationArea, err := cfg.pokeapiClient.GetLocationList(url)
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

func commandMapb(cfg *config) error {

	url := cfg.prevLocationsURL
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationArea, err := cfg.pokeapiClient.GetLocationList(url)
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
