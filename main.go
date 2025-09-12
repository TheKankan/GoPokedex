package main

import (
	"strings"
)

func main() {
	println("Hello, World!")
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	result := []string{}
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return result
}
