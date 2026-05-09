package main

import (
	"time"

	"github.com/m-heisters/pokedex-cli/internal/pokeapi"
	"github.com/m-heisters/pokedex-cli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokecache.NewCache(5 * time.Second),
	}
	startRepl(cfg)
}
