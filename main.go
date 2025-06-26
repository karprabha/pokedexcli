package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}

func main() {
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

		cmd, ok := commands[firstWord]
		if !ok {
			fmt.Printf("Unknown command: %s\n", firstWord)
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}
