package gosh

import (
	"unicode/utf8"
	"strconv"
	"github.com/scrouthtv/termios"
)

type Prompt struct {
	currentLine []rune
	parent      *Gosh
	promptStyle PromptStyle
	cursorPosition int
}

type PromptStyle interface {
	// LeftPrompt returns the contents of the area to the left of the prompt.
	// The int parameter indicates which prompt line is requested (starting with 0).
	// All paddings and margins must be added inside the prompt style, as the style will
	// be printed as-is and with no added spaces.
	LeftPrompt(*Gosh, int) string
	// RightPrompt returns the contents of the area on the right side of the screen.
	// The contents should be returned from left to right
	RightPrompt(*Gosh, int) string
}

func NewPrompt(parent *Gosh) *Prompt {
	return &Prompt{nil, parent, &SimplePromptStyle{}, 0}
}

func (p *Prompt) OnKey(key termios.Key) {
	if key == termios.InvalidKey {
		p.redraw()
	} else if key.Type == termios.KeyLetter {
		p.currentLine = append(p.currentLine, key.Value)
		p.cursorPosition += 1
	} else if key.Type == termios.KeySpecial {
		switch key.Value {
		case termios.SpecialBackspace:
			if len(p.currentLine) > 0 {
				p.cursorPosition -= 1
				p.currentLine = p.currentLine[:len(p.currentLine)-1]
			}
		case termios.SpecialEnter:
			p.cursorPosition = 0
			p.currentLine = nil
			p.parent.WriteString("\r\n")
		}
	}

	p.redraw()
}

func (p *Prompt) redraw() {
	p.parent.WriteString("\033[0G") // move to column 0
	p.parent.WriteString("\033[0J") // clear to end of screen, TODO: replace with termios.Action

	var leftPrompt string = p.promptStyle.LeftPrompt(p.parent, 0)
	p.parent.WriteString(leftPrompt)
	
	p.parent.WriteString(string(p.currentLine))

	var rightPrompt string = p.promptStyle.RightPrompt(p.parent, 0)
	p.parent.WriteString("\033[50G") // move to column 50, TODO move it so the text is printed to the right border
	p.parent.WriteString(rightPrompt)

	var position int = utf8.RuneCountInString(leftPrompt) + p.cursorPosition + 1
	p.parent.WriteString("\033[" + strconv.Itoa(position) + "G") // set cursor position
}
