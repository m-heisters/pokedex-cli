package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	baseXp := pokemon.BaseExperience
	rand := rand.Intn(99)

	if rand > baseXp {
		fmt.Printf("%s was caught!\n", name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	fmt.Println("Already caught: ")
	for key, _ := range cfg.caughtPokemon {
		fmt.Println(key)
	}

	return nil
}
