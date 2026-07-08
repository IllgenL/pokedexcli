package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/illgenl/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		text = strings.TrimSpace(text)
		words := cleanInput(text)
		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	cleanedWords := make([]string, len(words))
	for i, word := range words {
		cleanedWords[i] = strings.ToLower(word)
	}
	return cleanedWords
}
