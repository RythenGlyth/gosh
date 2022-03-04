package gosh

import (
	"gosh/src/shared"
	"os"
	"strings"
)

// SimplePromptStyle is a prompt style that displays the most basic contents
type SimplePromptStyle struct{}

// LeftPrompt displays the current working directory and an >
func (s *SimplePromptStyle) LeftPrompt(g shared.IGosh, line int) string {
	var prompt strings.Builder

	prompt.WriteString(" ")
	prompt.WriteString(s.pos(g))
	prompt.WriteString(" > ")

	return prompt.String()
}

func (s *SimplePromptStyle) pos(g shared.IGosh) string {
	path, err := g.GetWD()
	if err != nil {
		return "!"
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	if path == home {
		return "~"
	}

	return path
}

// RightPrompt simply displays a bar
func (s *SimplePromptStyle) RightPrompt(g shared.IGosh, line int) string {
	return " | "
}
