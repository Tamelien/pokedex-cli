package main

type cliCommand struct {
	name        string
	description string
	mapCommand  string
	callback    func(cfg *config, args ...string) error
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the last 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the Pokemons in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the name, height, weight, stats and type(s) of a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}
