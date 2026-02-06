package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("No location provided")
	}

	LocationEncounters, err := cfg.pokeapiClient.GetLocationEncounters(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\nFound Pokemon:\n", args[0])
	for _, encounter := range LocationEncounters.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
