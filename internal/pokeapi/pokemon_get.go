package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespPokemon{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return locationResp, nil
	}
	pokemon := RespPokemon{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemon, nil
}
