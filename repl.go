package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
