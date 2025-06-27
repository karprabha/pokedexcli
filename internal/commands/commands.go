package commands

import (
	"github.com/karprabha/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	Cache               *pokecache.Cache
	Pokedex             map[string]Pokemon
}

type Command struct {
	Name        string
	Description string
	callback    func(*Config, []string) error
}

func (c Command) Execute(cfg *Config, args []string) error {
	return c.callback(cfg, args)
}

var commands map[string]Command

func init() {
	commands = map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location area to see all Pokemon found there.",
			callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a Pokemon and add it to your Pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "View details about a caught Pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See all the pokemon in your Pokedex.",
			callback:    commandPokedex,
		},
	}
}

func GetCommand(name string) (Command, bool) {
	cmd, exists := commands[name]
	return cmd, exists
}

func GetAllCommands() map[string]Command {
	return commands
}
