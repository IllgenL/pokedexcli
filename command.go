package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays the encounterable pokemon in the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Displays pokemon information if pokemon was caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all your captured pokemon",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Prev

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Prev

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	encountersResp, err := cfg.pokeapiClient.ListEncounters(cfg.prevLocationsURL, name)
	if err != nil {
		return err
	}

	for _, encounter := range encountersResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	catchSuccess, pokemon, err := cfg.pokeapiClient.AttemptCatch(cfg.prevLocationsURL, name)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at", pokemon.Name+"...")
	if catchSuccess {
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Println(pokemon.Name, "was caught!")
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you should provide a pokemon name")
	}

	name := args[0]

	pokemon, ok := cfg.pokedex[name]

	if !ok {
		return errors.New(name + " was not found")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	fmt.Println("- hp:", pokemon.Stats[0].BaseStat)
	fmt.Println("- attack:", pokemon.Stats[1].BaseStat)
	fmt.Println("- defense:", pokemon.Stats[2].BaseStat)
	fmt.Println("- special attack:", pokemon.Stats[3].BaseStat)
	fmt.Println("- special defense:", pokemon.Stats[4].BaseStat)
	fmt.Println("- speed:", pokemon.Stats[5].BaseStat)
	if len(pokemon.Types) == 1 {
		fmt.Println("Type: ", pokemon.Types[0].Type.Name)
	} else {
		fmt.Println("Types:")
		fmt.Println("-", pokemon.Types[0].Type.Name)
		fmt.Println("-", pokemon.Types[1].Type.Name)
	}
	if len(pokemon.Abilities) == 1 {
		fmt.Println("Ability:", pokemon.Abilities[0].Ability.Name)
	} else {
		fmt.Println("Abilities:")
		for _, ability := range pokemon.Abilities {
			if ability.IsHidden {
				fmt.Println("-", ability.Ability.Name, "(Hidden Ability)")
			} else {
				fmt.Println("-", ability.Ability.Name)
			}
		}
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("your Pokedex is empty")
	}

	fmt.Println("Your Pokedex:")

	for name := range cfg.pokedex {
		fmt.Println("-", name)
	}

	return nil
}
