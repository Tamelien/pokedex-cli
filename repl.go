package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		arg := ""
		if len(input) > 1 {
			arg = input[1]
		}

		command, ok := getCommand()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, arg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(raw string) []string {
	text_lower := strings.ToLower(raw)
	return strings.Fields(text_lower)
}
