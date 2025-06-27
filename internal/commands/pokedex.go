package commands

import (
	"fmt"
)

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Printf("Your Pokedex is empty.\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")
	for name := range cfg.Pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
