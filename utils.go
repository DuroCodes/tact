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

func generateWords(count int, wordset string) []string {
	fallbackWords := []string{
		"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog",
	}

	file, err := os.Open("wordsets/" + wordset)
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

	result := make([]string, count)
	for i := 0; i < count; i++ {
		word := words[rand.Intn(len(words))]
		result[i] = strings.ToLower(word)
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

func contains(validWordsets []string, s string) bool {
	for _, wordset := range validWordsets {
		if wordset == s {
			return true
		}
	}
	return false
}
