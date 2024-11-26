package main

import (
	"flag"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	wordCount := flag.Int("n", 25, "number of words to type (min: 1, max: 100)")
	flag.Parse()

	if *wordCount < 1 {
		*wordCount = 1
	} else if *wordCount > 100 {
		*wordCount = 100
	}

	clearScreen()
	p := tea.NewProgram(initialModel(*wordCount))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
