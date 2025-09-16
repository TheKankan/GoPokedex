package main

import (
	"errors"
	"fmt"
	"math/rand"
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
		return errors.New("Please specify the zone you want to explore : explore [zone]")
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

func commandCatch(cfg *Config, parameter string) error {

	if parameter == "" {
		return errors.New("Please specify the pokemon you want to catch : catch [pokemon]")
	}
	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/" + parameter

	if _, exists := cfg.caughtPokemon[parameter]; exists {
		fmt.Println(parameter, "is already caught!")
		return nil
	}

	result, err := pokeapi.PokemonInfo(pokemonUrl, &cfg.pokeCache)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + parameter + "...")
	n := rand.Intn(1001)
	if n < result.BaseExperience {
		fmt.Println(parameter + " escaped!")
	} else {
		fmt.Println(parameter + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[parameter] = result
	}
	return nil
}

func commandInspect(cfg *Config, parameter string) error {

	if parameter == "" {
		return errors.New("Please specify the pokemon you want to inspect : inspect [pokemon]")
	}

	if _, exists := cfg.caughtPokemon[parameter]; exists {
		pokemon := cfg.caughtPokemon[parameter]
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Base Experience:", pokemon.BaseExperience)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Println(" -", t.Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

func commandPokedex(cfg *Config, parameter string) error {
	if len(cfg.caughtPokemon) != 0 {
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Println(" -", pokemon.Name)
		}
		return nil
	} else {
		fmt.Println("You haven't caught a Pokemon yet ! Use the catch command.")
		return nil
	}
}
