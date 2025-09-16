package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/TheKankan/GoPokedex/internal/pokecache"
)

func ListLocation(url string, cache *pokecache.Cache) (LocationList, error) {

	if val, ok := cache.Get(url); ok {
		var loc LocationList
		if err := json.Unmarshal(val, &loc); err != nil {
			return LocationList{}, err
		}
		return loc, nil
	}
	// not in cache, fetch from API

	res, err := http.Get(url)
	if err != nil {
		return LocationList{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationList{}, errors.New("error: received non-200 response code")
	}

	rawJson, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{}, err
	}

	var loc LocationList
	if err := json.Unmarshal(rawJson, &loc); err != nil {
		return LocationList{}, err
	}

	cache.Add(url, rawJson)
	return loc, nil

}

type LocationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
