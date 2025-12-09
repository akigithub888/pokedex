package main

import (
	"fmt"
	"os"
	"strings"
)

var Commands map[string]cliCommand

func cleanInput(input string) []string {
	// Trim leading/trailing spaces and set to lowercase
	clean := strings.ToLower(strings.TrimSpace(input))

	// Split by ANY amount of whitespace
	words := strings.Fields(clean)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, cmd := range Commands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func init() {
	Commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show all commands",
			callback:    commandHelp,
		},
	}
}
