package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TheKankan/GoPokedex/internal/pokecache"
)

func ListZonePokemon(url string, cache *pokecache.Cache) (PokemonList, error) {

	if val, ok := cache.Get(url); ok {
		var loc PokemonList
		if err := json.Unmarshal(val, &loc); err != nil {
			return PokemonList{}, err
		}
		return loc, nil
	}
	// not in cache, fetch from API

	res, err := http.Get(url)
	if err != nil {
		return PokemonList{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return PokemonList{}, errors.New("error: received non-200 response code")
	}

	rawJson, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonList{}, err
	}

	var pokemon PokemonList
	if err := json.Unmarshal(rawJson, &pokemon); err != nil {
		return PokemonList{}, err
	}

	cache.Add(url, rawJson)
	return pokemon, nil

}

type PokemonList struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			//URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
