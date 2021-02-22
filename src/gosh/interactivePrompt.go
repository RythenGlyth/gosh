package gosh

import (
	"container/list"
	"gosh/util"
	"strconv"
	"unicode/utf8"

	"github.com/scrouthtv/termios"
)

// Prompt is responsible for drawing the UI of an interactive Gosh
type Prompt struct {
	line   *list.List
	parent *Gosh
	style  PromptStyle
	pos    *list.Element // points to the element to the left of the cursor
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
	return &Prompt{list.New(), parent, &SimplePromptStyle{}, nil}
}

func (p *Prompt) doBackspace() {
	if p.line.Len() == 0 {
		return
	}

	var newPos *list.Element = p.pos.Prev()
	p.line.Remove(p.pos)

	if p.line.Len() == 0 {
		p.pos = nil
	} else {
		p.pos = newPos
	}
}

// OnKey is the event callback when a key has been pressed in the interactive terminal
func (p *Prompt) OnKey(key termios.Key) {
	if key == termios.InvalidKey {
		p.redraw()
	} else if key.Type == termios.KeyLetter {
		if p.pos == nil {
			p.line.PushBack(key.Value)
			p.pos = p.line.Back()
		} else {
			p.line.InsertAfter(key.Value, p.pos)
			p.pos = p.pos.Next()
		}
	} else if key.Type == termios.KeySpecial {
		switch key.Value {
		case termios.SpecialBackspace:
			p.doBackspace()
		case termios.SpecialEnter:
			var line string = util.RuneListToString(p.line)

			p.pos = nil
			p.line = list.New()
			p.parent.WriteString("\r\n")

			p.parent.Eval(line)
		case termios.SpecialArrowLeft:
			if p.pos != p.line.Front() {
				p.pos = p.pos.Prev()
			}
		case termios.SpecialArrowRight:
			if p.pos != p.line.Back() {
				p.pos = p.pos.Next()
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

	var line string = util.RuneListToString(p.line)
	var pos int = util.PositionInList(p.line, p.pos)
	// theoretically this could be done in a single loop, but I don't want to
	if pos == -1 && p.line.Len() > 0 {
		pos = 0
	}

	p.parent.WriteString(line)

	var rightPrompt string = p.style.RightPrompt(p.parent, 0)
	var width uint16 = p.parent.term.GetSize().Width
	if width == 0 {
		width = 80
	}
	var position int = int(width) - utf8.RuneCountInString(rightPrompt)
	p.parent.WriteString("\033[" + strconv.Itoa(position) + "G") // move to column 50, TODO move it so the text is printed to the right border
	p.parent.WriteString(rightPrompt)

	position = utf8.RuneCountInString(leftPrompt) + pos + 2
	p.parent.WriteString("\033[" + strconv.Itoa(position) + "G") // set cursor position
}
