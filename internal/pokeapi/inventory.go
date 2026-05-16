package pokeapi

import (
	"fmt"
	"sync"
)

type Inventory struct {
	inventory map[string]Pokemon
	mux       *sync.RWMutex
}

func NewInventory() *Inventory {
	return &Inventory{
		inventory: make(map[string]Pokemon),
		mux:       &sync.RWMutex{},
	}

}

func (i *Inventory) Add(pokemon Pokemon) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.inventory[pokemon.Name] = pokemon
}

func (i *Inventory) ListInventory() {
	for key, _ := range i.inventory {
		fmt.Println(key)
	}
}
