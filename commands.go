package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var config Config

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *Config
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      &config,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
			config:      &config,
		},
		"map": {
			name:        "map",
			description: "Displays name of next 20 locations",
			callback:    commandMap,
			config:      &config,
		},
		"mapb": {
			name:        "mapBack",
			description: "Displays name of previous 20 locations",
			callback:    commandMapBack,
			config:      &config,
		},
	}
	config = Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
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

func getLocationAreas(err error, res *http.Response) (location_areas_response, error) {

	location_areas := location_areas_response{}
	if err != nil {
		return location_areas, fmt.Errorf("Error calling PokeApi location-area endpoint")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return location_areas, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode > 299 {
		return location_areas, fmt.Errorf("PokeApi returned error code %d: %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, &location_areas)
	if err != nil {
		return location_areas, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return location_areas, nil
}

func commandMap() error {
	res, err := http.Get(config.Next)
	locationAreas, err := getLocationAreas(err, res)

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	if err != nil {
		return err
	}
	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	return nil
}

func commandMapBack() error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(config.Previous)
	locationAreas, err := getLocationAreas(err, res)

	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	if err != nil {
		return err
	}
	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous

	return nil
}
