package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	NextLocation     string
	PreviousLocation any
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

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
		cmdName := words[0]
		command, exists := getCommands()[cmdName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command : %s \n\n", cmdName)
			continue
		}

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

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next areas of the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous areas of the Pokemon world",
			callback:    commandMapB,
		},
	}
}
