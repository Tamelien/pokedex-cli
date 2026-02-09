package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("No pokemon to catch")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	if rand.Intn(pokemon.BaseExperience) >= 40 {
		fmt.Printf("%s escaped!\n", args[0])
	} else {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.caughtPokemon[args[0]] = pokemon
	}

	return nil
}
