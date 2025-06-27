package commands

import (
	"fmt"
)

func commandInspect(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: inspect <pokemon_name>")
	}

	pokemonName := args[0]
	pokemon, caught := cfg.Pokedex[pokemonName]
	if !caught {
		fmt.Printf("You have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
