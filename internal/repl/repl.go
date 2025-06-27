package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/karprabha/pokedexcli/internal/commands"
	"github.com/karprabha/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func Start() {
	cfg := &commands.Config{
		Cache:   pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]commands.Pokemon),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		firstWord := cleaned[0]
		args := cleaned[1:]

		cmd, ok := commands.GetCommand(firstWord)
		if !ok {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}

		err := cmd.Execute(cfg, args)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}
