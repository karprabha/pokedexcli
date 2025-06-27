package commands

import (
	"github.com/karprabha/pokedexcli/internal/pokecache"
)

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	Cache               *pokecache.Cache
}

type Command struct {
	Name        string
	Description string
	callback    func(*Config) error
}

func (c Command) Execute(cfg *Config) error {
	return c.callback(cfg)
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
	}
}

func GetCommand(name string) (Command, bool) {
	cmd, exists := commands[name]
	return cmd, exists
}

func GetAllCommands() map[string]Command {
	return commands
}
