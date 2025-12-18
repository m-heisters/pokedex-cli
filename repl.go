package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(strings.ToLower(text))
	return strings.Split(trimmed, " ")
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		userCommand := cleanInput(input)[0]
		command, ok := supportedCommands[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println("Something went wrong")
		}
	}

}
