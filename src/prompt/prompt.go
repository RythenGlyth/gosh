package prompt

import (
	"io"

	"github.com/scrouthtv/termios/utf8"
)

type Prompt struct {
	currentLine []byte
	currentSize int
	out io.Writer
}

func NewPrompt(out io.Writer) (*Prompt) {
	return &Prompt{nil, 1, out}
}

func (p *Prompt) OnKey(key utf8.Key) {
	switch key.Type {
	case utf8.KeyLetter:
		p.currentLine = append(p.currentLine, key.Value)
	case utf8.KeyCtrl:
		// TODO
	case utf8.KeyAlt:
		// TODO
	case utf8.KeyUml:
		p.currentLine = append(p.currentLine, 0xC3, key.Value)
	case utf8.KeySpecial:
		switch key.Value {
		case utf8.SpecialBackspace:
			p.currentLine = p.currentLine[:len(p.currentLine)-1]
		}
	}

	p.redraw()
}

func (p *Prompt) redraw() {
	p.out.Write([]byte("\033[0G")) // move to column 0
	p.out.Write([]byte("\033[0J")) // clear to end of screen
	p.out.Write(p.currentLine)
}
