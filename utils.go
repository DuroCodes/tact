package main

import (
	"bufio"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateWords(_count int) []string {
	fallbackWords := []string{
		"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog",
	}

	file, err := os.Open("words.txt")
	if err != nil {
		return fallbackWords
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if len(words) == 0 || scanner.Err() != nil {
		return fallbackWords
	}

	result := make([]string, _count)
	for i := 0; i < _count; i++ {
		result[i] = words[rand.Intn(len(words))]
	}
	return result
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
