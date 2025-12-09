package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		command := words[0]
		cmd, exists := Commands[command]
		if !exists {
			fmt.Println("Unknown command:", command)
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
