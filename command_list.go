package main

import (
	"fmt"
	"os"

	"github.com/TheKankan/GoPokedex/internal/pokeapi"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("  %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.NextLocation == "" {
		cfg.NextLocation = "https://pokeapi.co/api/v2/location-area/"
	}

	locations, err := pokeapi.ListLocation(&cfg.NextLocation)
	if err != nil {
		return err
	}

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	cfg.PreviousLocation = locations.Previous
	cfg.NextLocation = locations.Next
	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.PreviousLocation == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.ListLocation(&cfg.PreviousLocation)
	if err != nil {
		return err
	}

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	cfg.PreviousLocation = locations.Previous
	cfg.NextLocation = locations.Next
	return nil
}
