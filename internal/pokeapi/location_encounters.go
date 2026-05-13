package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(locationName string) (RespPokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName

	encounterResp := RespPokemonEncounters{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonEncounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	err = json.Unmarshal(data, &encounterResp)
	if err != nil {
		return RespPokemonEncounters{}, err
	}

	return encounterResp, nil
}
