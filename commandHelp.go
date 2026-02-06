package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, command := range getCommand() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
