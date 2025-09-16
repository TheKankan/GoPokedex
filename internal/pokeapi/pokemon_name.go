package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TheKankan/GoPokedex/internal/pokecache"
)

func PokemonInfo(url string, cache *pokecache.Cache) (PokemonInfos, error) {

	if val, ok := cache.Get(url); ok {
		var info PokemonInfos
		if err := json.Unmarshal(val, &info); err != nil {
			return PokemonInfos{}, err
		}
		return info, nil
	}
	// not in cache, fetch from API

	res, err := http.Get(url)
	if err != nil {
		return PokemonInfos{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PokemonInfos{}, errors.New("pokemon not found")
	}

	if res.StatusCode > 299 {
		return PokemonInfos{}, errors.New("error: received non-200 response code")
	}

	rawJson, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfos{}, err
	}

	var pokemon PokemonInfos
	if err := json.Unmarshal(rawJson, &pokemon); err != nil {
		return PokemonInfos{}, err
	}

	cache.Add(url, rawJson)
	return pokemon, nil

}

type PokemonInfos struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
