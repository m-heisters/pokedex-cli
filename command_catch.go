package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	encounterResponse, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	fmt.Println("Found Pokemon: ")

	for _, encounter := range encounterResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
