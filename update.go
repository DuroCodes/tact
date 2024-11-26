package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		m.wpmTickCmd(),
	)
}

func (m model) wpmTickCmd() tea.Cmd {
	return tea.Every(250*time.Millisecond, func(t time.Time) tea.Msg {
		return wpmTickMsg(t)
	})
}

type wpmTickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case wpmTickMsg:
		if m.hasStarted && !m.finished {
			duration := time.Since(m.startTime)
			m.lastWPM = float64(m.correctChars) / 5.0 / duration.Minutes()
		}
		cmds = append(cmds, m.wpmTickCmd())
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			return initialModel(m.wordCount), nil
		case tea.KeySpace:
			if !m.hasStarted {
				m.hasStarted = true
				m.startTime = time.Now()
				cmds = append(cmds, m.stopwatch.Init())
			}
			if m.currentWord < len(m.words)-1 {
				for len(m.typedWords) <= m.currentWord {
					m.typedWords = append(m.typedWords, "")
				}
				m.typedWords[m.currentWord] = m.inputText

				m.correctChars = 0
				m.totalChars = 0
				for i, typed := range m.typedWords {
					if i >= len(m.words) {
						break
					}
					m.totalChars += len(m.words[i])
					if typed == m.words[i] {
						m.correctChars += len(m.words[i])
					}
				}

				m.currentWord++
				m.inputText = ""
			}
		case tea.KeyBackspace:
			if !m.hasStarted {
				m.hasStarted = true
				m.startTime = time.Now()
				cmds = append(cmds, m.stopwatch.Init())
			}
			if len(m.inputText) > 0 {
				if msg.Alt {
					lastSpace := strings.LastIndex(m.inputText[:len(m.inputText)-1], " ")
					if lastSpace >= 0 {
						m.inputText = m.inputText[:lastSpace+1]
					} else {
						m.inputText = ""
					}
				} else {
					m.inputText = m.inputText[:len(m.inputText)-1]
				}
			} else if m.currentWord > 0 {
				m.currentWord--
				m.inputText = m.typedWords[m.currentWord]
			}
		default:
			if !m.finished {
				if !m.hasStarted {
					m.hasStarted = true
					m.startTime = time.Now()
					cmds = append(cmds, m.stopwatch.Init())
				}

				if len(m.inputText) < len(m.words[m.currentWord]) {
					m.inputText += msg.String()
				}

				if m.currentWord == len(m.words)-1 && m.inputText == m.words[m.currentWord] {
					if m.inputText == m.words[m.currentWord] {
						m.correctChars += len(m.words[m.currentWord])
					}
					m.totalChars += len(m.words[m.currentWord])
					m.finished = true
					m.endTime = time.Now()
				}
			}
		}
	}

	m.stopwatch, cmd = m.stopwatch.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
