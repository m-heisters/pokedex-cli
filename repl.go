package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m-heisters/pokedex-cli/internal/pokeapi"
	"github.com/m-heisters/pokedex-cli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	cache            *pokecache.Cache
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	return strings.Split(trimmed, " ")
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		parts := cleanInput(scanner.Text())
		commandName := parts[0]
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, arg)
		if err != nil {
			fmt.Println("Something went wrong")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area and discover encounters",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
