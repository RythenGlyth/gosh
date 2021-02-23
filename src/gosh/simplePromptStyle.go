package gosh

import "strings"

// SimplePromptStyle is a prompt style that displays the most basic contents
type SimplePromptStyle struct{}

// LeftPrompt displays the current working directory and an >
func (s *SimplePromptStyle) LeftPrompt(g *Gosh, line int) string {
	var prompt strings.Builder

	prompt.WriteString(" ")
	prompt.WriteString(g.GetWD())
	prompt.WriteString(" > ")

	return prompt.String()
}

// RightPrompt simply displays a bar
func (s *SimplePromptStyle) RightPrompt(g *Gosh, line int) string {
	return " | "
}
