package main

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	textInput    textinput.Model
	words        []string
	currentWord  int
	inputText    string
	correctChars int
	totalChars   int
	startTime    time.Time
	finished     bool
	hasStarted   bool
	endTime      time.Time
	lastWPM      float64
	typedWords   []string
	progress     progress.Model
	stopwatch    stopwatch.Model
	wpmTicker    *time.Ticker
	wordCount    int
	wordset      string
}

func initialModel(wordCount int, wordset string) model {
	ti := textinput.New()
	ti.Focus()

	prog := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(80),
	)

	return model{
		textInput:  ti,
		words:      generateWords(wordCount, wordset),
		startTime:  time.Now(),
		hasStarted: false,
		typedWords: make([]string, 0),
		progress:   prog,
		stopwatch:  stopwatch.NewWithInterval(time.Millisecond),
		wpmTicker:  time.NewTicker(250 * time.Millisecond),
		wordCount:  wordCount,
	}
}
