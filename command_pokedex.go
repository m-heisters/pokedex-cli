package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your pokedex: ")
	for key := range cfg.caughtPokemon {
		fmt.Println(key)
	}
	return nil
}
