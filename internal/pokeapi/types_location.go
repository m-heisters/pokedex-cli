package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonEncounters struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Results []struct {
		Name string `json:"name"`
	} `json:"Pokemon_encounters"`
}
