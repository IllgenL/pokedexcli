package pokeapi

type RespShallowLocations struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"prev"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
