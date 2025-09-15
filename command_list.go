package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
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

	res, err := http.Get(cfg.NextLocation)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("error: received non-200 response code")
	}

	var loc LocationList
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return err
	}

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}
	cfg.PreviousLocation = loc.Previous
	cfg.NextLocation = loc.Next
	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.PreviousLocation == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(cfg.PreviousLocation.(string))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("error: received non-200 response code")
	}

	var loc LocationList
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return err
	}

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}
	cfg.PreviousLocation = loc.Previous
	cfg.NextLocation = loc.Next
	return nil
}

type LocationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
