package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	var err error
	if pageURL != nil {
		url = *pageURL
	}

	locations := RespShallowLocations{}

	entry, ok := c.cache.Get(url)
	if ok {
		err = json.Unmarshal(entry, &locations)
		if err != nil {
			return locations, err
		}
		return locations, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	return locations, nil
}
