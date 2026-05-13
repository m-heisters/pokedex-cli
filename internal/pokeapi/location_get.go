package pokeapi

import (
	"encoding/json"
	"github.com/m-heisters/pokedex-cli/internal/pokecache"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string, cache *pokecache.Cache) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}
	locationResp := Location{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	return locationResp, nil
}
