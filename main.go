package main

import "github.com/akigithub888/pokedex/internal/pokeapi"

func main() {
	cfg := &Config{
		Client: *pokeapi.NewClient(),
	}
	startRepl(cfg)

}
