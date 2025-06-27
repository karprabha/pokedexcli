package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}

	pokemonName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	if _, caught := cfg.Pokedex[pokemonName]; caught {
		fmt.Printf("You have already caught %s!\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	var pokemon Pokemon
	if cachedData, found := cfg.Cache.Get(url); found {
		if err := json.Unmarshal(cachedData, &pokemon); err != nil {
			return err
		}
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Pokemon %s not found!\n", pokemonName)
			return nil
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get pokemon: %s", resp.Status)
		}

		cfg.Cache.Add(url, body)

		if err := json.Unmarshal(body, &pokemon); err != nil {
			return err
		}
	}

	// Calculate catch probability based on base experience
	// Higher base experience = harder to catch
	// Max base experience is typically around 600, so we'll use that as reference
	maxBaseExp := 600
	if pokemon.BaseExperience > maxBaseExp {
		maxBaseExp = pokemon.BaseExperience
	}

	// Catch probability: 50% base + up to 40% bonus for low base experience
	// This means legendary Pokemon (high base exp) are harder to catch
	catchProbability := 0.5 + (0.4 * float64(maxBaseExp-pokemon.BaseExperience) / float64(maxBaseExp))

	// Ensure probability is between 0.1 and 0.9
	if catchProbability < 0.1 {
		catchProbability = 0.1
	}
	if catchProbability > 0.9 {
		catchProbability = 0.9
	}

	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Printf("You may now inspect it with the inspect command.\n")

		cfg.Pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
