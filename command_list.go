package main

import (
	"fmt"
	"os"

	"github.com/TheKankan/GoPokedex/internal/pokeapi"
)

func commandExit(cfg *Config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, parameter string) error {
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

func commandMap(cfg *Config, parameter string) error {
	if cfg.NextLocation == "" {
		cfg.NextLocation = "https://pokeapi.co/api/v2/location-area/"
	}

	locations, err := pokeapi.ListLocation(cfg.NextLocation, &cfg.pokeCache)
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

func commandMapB(cfg *Config, parameter string) error {
	if cfg.PreviousLocation == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.ListLocation(cfg.PreviousLocation, &cfg.pokeCache)
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

func commandExplore(cfg *Config, parameter string) error {

	if parameter == "" {
		fmt.Println("Please specify the zone you want to explore : explore [zone]")
		return nil
	}
	exploreUrl := "https://pokeapi.co/api/v2/location-area/" + parameter

	result, err := pokeapi.ListZonePokemon(exploreUrl, &cfg.pokeCache)
	if err != nil {
		return err
	}

	for _, encounter := range result.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
