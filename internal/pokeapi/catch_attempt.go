package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) AttemptCatch(pageUrl *string, pokemonName string) (bool, Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	var err error
	if pageUrl != nil {
		return false, Pokemon{}, errors.New("missing page URL")
	}

	pokemon := Pokemon{}

	entry, ok := c.cache.Get(url)
	if ok {
		err = json.Unmarshal(entry, &pokemon)
		if err != nil {
			return false, Pokemon{}, err
		}
	} else {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return false, Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return false, Pokemon{}, err
		}
		defer res.Body.Close()

		dat, err := io.ReadAll(res.Body)
		if err != nil {
			return false, Pokemon{}, err
		}

		err = json.Unmarshal(dat, &pokemon)
		if err != nil {
			return false, Pokemon{}, err
		}
	}

	catch := rand.Intn(pokemon.BaseExperience)
	if catch > 40 {
		return false, pokemon, nil
	}

	return true, pokemon, nil
}
