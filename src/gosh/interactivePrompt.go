package gosh

import (
	"strconv"
	"unicode/utf8"

	"github.com/scrouthtv/termios"
)

// Prompt is responsible for drawing the UI of an interactive Gosh
type Prompt struct {
	line   []rune
	parent *Gosh
	style  PromptStyle
	pos    int
}

// PromptStyle are individual styles for the prompt.
// They create prompt text to the left and right.
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

// NewPrompt creates a new prompt with the SimplePromptStyle style.
func NewPrompt(parent *Gosh) *Prompt {
	return &Prompt{nil, parent, &SimplePromptStyle{}, 0}
}

func (p *Prompt) doBackspace() {
	if p.pos == 0 {
		return
	} else if p.pos == len(p.line)+1 {
		p.pos--
		p.line = p.line[:len(p.line)-1]
	} else {
		copy(p.line[p.pos:], p.line[p.pos+1:])
		p.pos--
	}
}

// OnKey is the event callback when a key has been pressed in the interactive terminal
func (p *Prompt) OnKey(key termios.Key) {
	if key == termios.InvalidKey {
		p.redraw()
	} else if key.Type == termios.KeyLetter {
		p.line = append(p.line, key.Value)
		p.pos++
	} else if key.Type == termios.KeySpecial {
		switch key.Value {
		case termios.SpecialBackspace:
			if len(p.line) > 0 {
				p.doBackspace()
			}
		case termios.SpecialEnter:
			p.pos = 0
			p.line = nil
			p.parent.WriteString("\r\n")
		case termios.SpecialArrowLeft:
			if p.pos > 0 {
				p.doBackspace()
			}
		case termios.SpecialArrowRight:
			if p.pos < len(p.line) {
				p.pos++
			}
		}
	}

	p.redraw()
}

func (p *Prompt) redraw() {
	p.parent.WriteString("\033[0G") // move to column 0
	p.parent.WriteString("\033[0J") // clear to end of screen, TODO: replace with termios.Action

	var leftPrompt string = p.style.LeftPrompt(p.parent, 0)
	p.parent.WriteString(leftPrompt)

	p.parent.WriteString(string(p.line))

	var rightPrompt string = p.style.RightPrompt(p.parent, 0)
	p.parent.WriteString("\033[50G") // move to column 50, TODO move it so the text is printed to the right border
	p.parent.WriteString(rightPrompt)

	var position int = utf8.RuneCountInString(leftPrompt) + p.pos + 1
	p.parent.WriteString("\033[" + strconv.Itoa(position) + "G") // set cursor position
}
