package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (m model) View() string {
	if m.finished {
		duration := m.endTime.Sub(m.startTime)
		wpm := float64(m.correctChars) / 5.0 / duration.Minutes()
		accuracy := float64(m.correctChars) / float64(m.totalChars) * 100

		titleStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7AA2F7")).
			Bold(true).
			Margin(1)

		statStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9ECE6A")).
			Padding(0, 2)

		instructionStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#565F89")).
			Margin(1)

		timeStr := ""
		if duration.Seconds() < 60 {
			timeStr = fmt.Sprintf("%.2fs", duration.Seconds())
		} else {
			timeStr = fmt.Sprintf("%d:%05.2f",
				int(duration.Minutes()),
				duration.Seconds()-float64(int(duration.Minutes())*60))
		}

		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s",
			titleStyle.Render("Test complete!"),
			statStyle.Render(fmt.Sprintf("WPM: %.2f", wpm)),
			statStyle.Render(fmt.Sprintf("Accuracy: %.2f%%", accuracy)),
			statStyle.Render(fmt.Sprintf("Time: %s", timeStr)),
			instructionStyle.Render("Press Tab to restart or Ctrl+C to quit"),
		)
	}

	var s strings.Builder
	maxWidth := 80

	m.progress.Width = maxWidth
	progress := float64(m.currentWord) / float64(len(m.words))
	progressBar := m.progress.ViewAs(progress)
	s.WriteString(progressBar + "\n\n")

	var lineContent strings.Builder
	normalStyle := lipgloss.NewStyle()
	wrongStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5C57")).Underline(true)
	cursorStyle := lipgloss.NewStyle().Reverse(true)
	futureStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6272A4"))

	for i, word := range m.words {
		if i == m.currentWord {
			for j, char := range word {
				if j < len(m.inputText) {
					if string(m.inputText[j]) == string(char) {
						lineContent.WriteString(normalStyle.Render(string(char)))
					} else {
						lineContent.WriteString(wrongStyle.Render(string(char)))
					}
				} else if j == len(m.inputText) {
					lineContent.WriteString(cursorStyle.Render(string(char)))
				} else {
					lineContent.WriteString(normalStyle.Render(string(char)))
				}
			}
		} else if i < m.currentWord {
			typedWord := m.typedWords[i]
			for j, char := range word {
				if j < len(typedWord) {
					if string(typedWord[j]) == string(char) {
						lineContent.WriteString(normalStyle.Render(string(char)))
					} else {
						lineContent.WriteString(wrongStyle.Render(string(char)))
					}
				} else {
					lineContent.WriteString(wrongStyle.Render(string(char)))
				}
			}
		} else {
			lineContent.WriteString(futureStyle.Render(word))
		}
		lineContent.WriteString(" ")
	}

	wrapped := wordwrap.String(lineContent.String(), maxWidth)
	s.WriteString(wrapped)

	statusStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
	wordsTyped := fmt.Sprintf("%d/%d", m.currentWord, len(m.words))
	accuracy := float64(m.correctChars) / float64(max(m.totalChars, 1)) * 100
	wpm := m.lastWPM
	if !m.hasStarted {
		wpm = 0
	}

	elapsed := "0.00"
	if m.hasStarted {
		elapsedTime := m.stopwatch.Elapsed()
		if elapsedTime.Seconds() < 60 {
			elapsed = fmt.Sprintf("%.2fs", elapsedTime.Seconds())
		} else {
			elapsed = fmt.Sprintf("%d:%05.2f",
				int(elapsedTime.Minutes()),
				elapsedTime.Seconds()-float64(int(elapsedTime.Minutes())*60))
		}
	}

	status := fmt.Sprintf("\n %s • %.0f • %.0f%% • %s",
		wordsTyped, wpm, accuracy, elapsed)
	s.WriteString(statusStyle.Render(status))

	s.WriteString("\n")
	return s.String()
}
