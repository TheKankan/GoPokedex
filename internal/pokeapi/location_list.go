package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ListLocation(url *string) (LocationList, error) {
	res, err := http.Get(*url)
	if err != nil {
		return LocationList{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationList{}, errors.New("error: received non-200 response code")
	}

	var loc LocationList
	if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
		return LocationList{}, err
	}
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
