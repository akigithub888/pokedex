package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akigithub888/pokedex/internal/pokeapi"
)

var Commands map[string]cliCommand

type Config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			return
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		cmd, ok := Commands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	// Trim leading/trailing spaces and set to lowercase
	clean := strings.ToLower(strings.TrimSpace(input))

	// Split by ANY amount of whitespace
	words := strings.Fields(clean)
	return words
}

func commandMap(cfg *Config) error {
	pageURL := ""
	if cfg.Next != nil {
		pageURL = *cfg.Next
	}

	list, err := cfg.Client.ListLocationAreas(pageURL)
	if err != nil {
		return err
	}

	cfg.Next = list.Next
	cfg.Previous = list.Previous

	for _, result := range list.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	pageURL := *cfg.Previous

	list, err := cfg.Client.ListLocationAreas(pageURL)
	if err != nil {
		return err
	}

	cfg.Next = list.Next
	cfg.Previous = list.Previous

	for _, result := range list.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
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
	callback    func(*Config) error
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
		"map": {
			name:        "map",
			description: "Show the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
