package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	encounterResponse, err := cfg.pokeapiClient.GetLocation(name, cfg.cache)
	if err != nil {
		return err
	}

	for _, encounter := range encounterResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon)
	}
	return nil
}
