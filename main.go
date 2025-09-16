package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/TheKankan/GoPokedex/internal/pokeapi"
	"github.com/TheKankan/GoPokedex/internal/pokecache"
)

type Config struct {
	pokeCache        pokecache.Cache
	NextLocation     string
	PreviousLocation string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		pokeCache: pokeapi.NewCache(5 * time.Second),
	}

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
		cmdParameter := ""
		if len(words) == 2 {
			cmdParameter = words[1]
		}
		command, exists := getCommands()[cmdName]
		if exists {
			err := command.callback(cfg, cmdParameter)
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
	callback    func(*Config, string) error
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
		"explore": {
			name:        "explore [zone]",
			description: "Explores the specified zone to find Pokemons",
			callback:    commandExplore,
		},
	}
}
