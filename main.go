package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	wordCount := flag.Int("n", 25, "number of words to type (min: 1, max: 100)")
	wordset := flag.String("w", "english", "word set to use")
	flag.Parse()

	if *wordCount < 1 {
		*wordCount = 1
	} else if *wordCount > 100 {
		*wordCount = 100
	}

	wordsDirEntries, err := os.ReadDir("wordsets")
	if err != nil {
		fmt.Printf("Error reading words directory: %v", err)
		os.Exit(1)
	}

	validWordsets := []string{}
	for _, entry := range wordsDirEntries {
		validWordsets = append(validWordsets, entry.Name())
	}

	if !contains(validWordsets, *wordset) {
		*wordset = "english"
	}

	clearScreen()
	p := tea.NewProgram(initialModel(*wordCount, *wordset))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
