package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const POKE_API_LOCATION_AREA = "https://pokeapi.co/api/v2/location-area/"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays name of 20 locations",
			callback:    commandMap,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, item := range supportedCommands {
		fmt.Printf("%v: %v\n", item.name, item.description)

	}
	return nil
}

func commandMap() error {
	res, err := http.Get(POKE_API_LOCATION_AREA)
	if err != nil {
		fmt.Println("Error calling PokeApi location-area endpoint")
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Und dann separat diese:
	if res.StatusCode > 299 {
		return fmt.Errorf("PokeApi returned error code %d: %s", res.StatusCode, body)
	}

	location_area := location_areas{}
	err = json.Unmarshal(body, &location_area)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	fmt.Println(location_area)
	return nil
}
