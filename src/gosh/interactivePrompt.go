package gosh

import (
	"container/list"
	"gosh/src/shared"
	"gosh/src/util"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/scrouthtv/termios"
)

// Prompt is responsible for drawing the UI of an interactive Gosh
type Prompt struct {
	line   *list.List
	parent *Gosh
	style  PromptStyle
	pos    *list.Element // points to the element to the left of the cursor
	last   []string
	lastY  int
}

// PromptStyle are individual styles for the prompt.
// They create prompt text to the left and right.
type PromptStyle interface {

	// LeftPrompt returns the contents of the area to the left of the prompt.
	// The int parameter indicates which prompt line is requested (starting with 0).
	// All paddings and margins must be added inside the prompt style, as the style will
	// be printed as-is and with no added spaces.
	LeftPrompt(shared.IGosh, int) string

	// RightPrompt returns the contents of the area on the right side of the screen.
	// The contents should be returned from left to right
	RightPrompt(shared.IGosh, int) string
}

// NewPrompt creates a new prompt with the SimplePromptStyle style.
func NewPrompt(parent *Gosh) *Prompt {
	return &Prompt{list.New(), parent, &SimplePromptStyle{}, nil, nil, 0}
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
			p.last = nil
			p.lastY = 0
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

// lines returns
//  - the lines that should be drawn
//  - the x & y cursor position from the top left
func (p *Prompt) lines() ([]string, int, int) {
	var lines []string

	width := int(p.parent.term.GetSize().Width)

	var line strings.Builder

	leftPrompt := p.style.LeftPrompt(p.parent, 0)
	line.WriteString(leftPrompt)
	line.WriteString(util.RuneListToString(p.line))

	rightPrompt := p.style.RightPrompt(p.parent, 0)
	spaces := width - line.Len() - utf8.RuneCountInString(rightPrompt)

	for i := 0; i < spaces; i++ {
		line.WriteString(" ")
	}

	line.WriteString(rightPrompt)

	lines = append(lines, line.String())

	var xPos int = util.PositionInList(p.line, p.pos)
	if xPos == -1 && p.line.Len() > 0 {
		xPos = 0
	}
	xPos += utf8.RuneCountInString(leftPrompt) + 1

	return lines, xPos, 0
}

func (p *Prompt) redraw() {
	// hack time: hide the cursor while redrawing:
	p.parent.WriteString("\x1b[?25l")

	if p.lastY < 1 {
		p.parent.WriteString("\r") // move to the beginning of this line
	} else {
		p.parent.WriteString("\r\x1b[" + strconv.Itoa(p.lastY-1) + "A") // move to the beginning n lines up
	}

	var lines []string
	var xPos, yPos int
	lines, xPos, yPos = p.lines()

	var end1 int = len(lines)
	if len(p.last) < end1 {
		end1 = len(p.last)
	}

	var i int
	for i = 0; i < end1; i++ {
		if lines[i] != p.last[i] {
			p.parent.WriteString(lines[i])
		}
		p.parent.WriteString("\r\n")
	}

	// draw all added lines:
	for ; i < len(lines); i++ {
		p.parent.WriteString(lines[i])
		p.parent.WriteString("\r\n")
	}

	// place the cursor:
	var up int = len(lines) - yPos
	if up > 0 {
		p.parent.WriteString("\x1b[" + strconv.Itoa(up) + "A") // up
	}
	p.parent.WriteString("\r\x1b[" + strconv.Itoa(xPos) + "C")

	p.parent.WriteString("\x1b[?25h")

	p.last = lines
	p.lastY = yPos
}
