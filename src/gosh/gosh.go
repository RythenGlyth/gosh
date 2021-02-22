package gosh

import (
	"fmt"
	"os"

	"github.com/scrouthtv/termios"
)

// Gosh type collects all modules of a gosh shell.
type Gosh struct {
	term   termios.Terminal
	prompt *Prompt
	ready  bool
}

// NewGosh creates a new, empty gosh.
func NewGosh() *Gosh {
	return &Gosh{nil, nil, false}
}

// Init prepares all sub-functionality of this gosh instance.
// If the returned error is not nil, this gosh must not be used until
// the issue is fixed.
func (g *Gosh) Init() error {
	var err error
	var term termios.Terminal
	term, err = termios.Open()
	if err != nil {
		return err
	}
	g.term = term

	g.prompt = NewPrompt(g)

	g.ready = true

	return nil
}

// Close closes all sub-parts of the gosh instance (I/O, the terminal).
func (g *Gosh) Close() {
	g.WriteString("\r\n")

	g.ready = false
	if g.term != nil {
		g.term.Close()
	}
}

var keyCd termios.Key = termios.Key{Type: termios.KeyLetter, Mod: termios.ModCtrl, Value: 'd'}

// Interactive sets this gosh's mode to interactive.
// User input is read from the underlying terminal
// and commands are executed in the current namespace.
// A suitable exit code and any error is returned in the end.
func (g *Gosh) Interactive() (int, error) {
	var err error

	if !g.ready {
		err = g.Init()
		if err != nil {
			os.Stdout.WriteString("Error occured during intialization:")
			os.Stdout.WriteString(err.Error())
			os.Stdout.WriteString("\n")
			return 1, err
		}
		defer g.Close()
	}

	var in []termios.Key
	var k termios.Key

	g.term.Write([]byte(fmt.Sprintf("This is %s %s. Press C-d to exit.\r\n", GoshName, GoshVersion)))
	g.prompt.redraw()

	for {
		in, err = g.term.Read()
		if err != nil {
			// Consider g.term broken:
			os.Stdout.WriteString("Error reading input:\n")
			os.Stdout.WriteString(err.Error() + "\n")
		} else {
			for _, k = range in {
				if k.Equal(&keyCd) { // C-d to quit
					return 0, nil
				} else {
					// TODO event handling
					//  -> plugin management
					g.prompt.OnKey(k)
				}
			}
		}
	}
}

func (g *Gosh) Write(p []byte) (int, error) {
	return g.term.Write(p)
}

// WriteString writes the specified string to the gosh's terminal
func (g *Gosh) WriteString(s string) (int, error) {
	return g.term.WriteString(s)
}
