package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaDetail struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <area_name>")
	}

	areaName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	if cachedData, found := cfg.Cache.Get(url); found {
		return processExploreResponse(cachedData)
	}

	fmt.Printf("Exploring %s...\n", areaName)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get location area: %s", resp.Status)
	}

	cfg.Cache.Add(url, body)

	return processExploreResponse(body)
}

func processExploreResponse(body []byte) error {
	locationDetail := LocationAreaDetail{}
	if err := json.Unmarshal(body, &locationDetail); err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
