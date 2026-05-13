package main

import (
	"fmt"
)

func commandExplore(cfg *config, locationName string) error {
	fmt.Println("CommmandExplore with locationName" + locationName)
	encounterResponse, err := cfg.pokeapiClient.ListEncounters(locationName)
	if err != nil {
		return err
	}

	for _, encounter := range encounterResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon)
	}
	return nil
}
