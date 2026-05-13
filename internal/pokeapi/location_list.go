package pokeapi

import (
	"encoding/json"
	"github.com/m-heisters/pokedex-cli/internal/pokecache"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	locationsResp := RespShallowLocations{}

	val, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(val, &locationsResp)
		if err == nil {
			return locationsResp, nil
		}

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp = RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	cache.Add(url, data)

	return locationsResp, nil
}
