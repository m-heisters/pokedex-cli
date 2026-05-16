package main

import (
	"time"

	"github.com/m-heisters/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	inventory := pokeapi.NewInventory()
	cfg := &config{
		pokeapiClient: pokeClient,
		Inventory:     inventory,
	}
	startRepl(cfg)
}
