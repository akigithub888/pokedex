package main

import "strings"

func cleanInput(input string) []string {
	// Trim leading/trailing spaces and set to lowercase
	clean := strings.ToLower(strings.TrimSpace(input))

	// Split by ANY amount of whitespace
	words := strings.Fields(clean)
	return words
}
