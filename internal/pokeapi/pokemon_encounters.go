package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(pageUrl *string, locationName string) (RespEncounterLocations, error) {
	url := baseURL + "/location-area/" + locationName
	var err error
	if pageUrl != nil {
		return RespEncounterLocations{}, nil
	}

	encounters := RespEncounterLocations{}

	entry, ok := c.cache.Get(url)
	if ok {
		err = json.Unmarshal(entry, &encounters)
		if err != nil {
			return encounters, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return encounters, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return encounters, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return encounters, err
	}

	err = json.Unmarshal(dat, &encounters)
	if err != nil {
		return encounters, err
	}

	c.cache.Add(url, dat)

	return encounters, nil
}
