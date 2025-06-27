package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {
	locationAreasURL := "https://pokeapi.co/api/v2/location-area"
	if cfg.NextLocationURL != nil {
		locationAreasURL = *cfg.NextLocationURL
	}

	if cachedData, found := cfg.Cache.Get(locationAreasURL); found {
		fmt.Println("Using cached data...")
		return processLocationResponse(cachedData, cfg)
	}

	fmt.Println("Making API request...")
	resp, err := http.Get(locationAreasURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get location areas: %s", resp.Status)
	}

	cfg.Cache.Add(locationAreasURL, body)

	return processLocationResponse(body, cfg)
}

func commandMapb(cfg *Config) error {
	if cfg.PreviousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	if cachedData, found := cfg.Cache.Get(*cfg.PreviousLocationURL); found {
		fmt.Println("Using cached data...")
		return processLocationResponse(cachedData, cfg)
	}

	fmt.Println("Making API request...")
	resp, err := http.Get(*cfg.PreviousLocationURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get location areas: %s", resp.Status)
	}

	cfg.Cache.Add(*cfg.PreviousLocationURL, body)

	return processLocationResponse(body, cfg)
}

func processLocationResponse(body []byte, cfg *Config) error {
	locationAreasResp := LocationAreasResp{}
	if err := json.Unmarshal(body, &locationAreasResp); err != nil {
		return err
	}

	cfg.NextLocationURL = locationAreasResp.Next
	cfg.PreviousLocationURL = locationAreasResp.Previous

	for _, area := range locationAreasResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}
