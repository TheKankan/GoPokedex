package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break // exit on EOF or error
		}
		line := scanner.Text() // get current line
		words := cleanInput(line)
		if len(words) == 0 {
			fmt.Printf("No input provided \n\n")
			continue // ignore empty input
		}
		fmt.Printf("Your command was: %s\n\n", words[0])
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)

	result := []string{}
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}
	return result
}
