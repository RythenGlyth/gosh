package gosh

type SimplePromptStyle struct {}

func (s *SimplePromptStyle) LeftPrompt(g *Gosh, line int) string {
	return " > "
}

func (s *SimplePromptStyle) RightPrompt(g *Gosh, line int) string {
	return " | "
}