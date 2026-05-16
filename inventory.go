package main

import (
	"fmt"
	"github.com/m-heisters/pokedex-cli/internal/pokeapi"
	"sync"
)

type Inventory struct {
	inventory map[string]pokeapi.Pokemon
	mux       *sync.RWMutex
}

func (i *Inventory) Add(name string, pokemon pokeapi.Pokemon) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.inventory[name] = pokemon
}

func (i *Inventory) ListInventory() {
	for key, value := range i.inventory {
		fmt.Println(key, value)
	}
}
