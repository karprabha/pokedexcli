package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/karprabha/pokedexcli/internal/commands"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func Start() {
	cfg := &commands.Config{}
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

		cmd, ok := commands.GetCommand(firstWord)
		if !ok {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}

		err := cmd.Execute(cfg)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}
