package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Inspecting %s...\n", args[0])

	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New("Pokemon not yet caught")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Heigth: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %v\n", stat.Stat.Name, stat.BaseStat)

	}
	fmt.Println("Types: ")

	for _, typ := range pokemon.Types {
		fmt.Printf("-%s\n", typ.Type.Name)
	}
	return nil

}
