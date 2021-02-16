package main

import (
	"os"
	"fmt"

	"gosh/prompt"
	"gosh/utf8"

	"github.com/scrouthtv/termios"
)

var MyGosh *Gosh

type Gosh struct {
	term termios.Terminal
	p *prompt.Prompt
	ready bool
}

func NewGosh() (*Gosh) {
	return &Gosh{nil, nil, false}
}

func (g *Gosh) Init() error {
	var err error
	var term termios.Terminal
	term, err = termios.Open()
	if err != nil {
		return err
	}
	err = term.SetRaw(true)
	if err != nil {
		return err
	}
	g.term = term

	g.p = prompt.NewPrompt(term)
	return nil
}

func (g *Gosh) Close() {
	g.ready = false
	if (g.term != nil) {
		g.term.Close()
	}
}

func main() {
	var retcode int = 0
	defer func() { os.Exit(retcode) }()

	MyGosh = NewGosh()

	var err error
	err = MyGosh.Init()
	if err != nil {
		os.Stdout.WriteString("Error occured during intialization:")
		os.Stdout.WriteString(err.Error())
		os.Stdout.WriteString("\n")
		retcode = 1
		return
	}
	defer MyGosh.Close()

	var n int
	var buf []byte = make([]byte, 1024)
	var key utf8.Key
	var keys []utf8.Key

	MyGosh.term.Write([]byte(fmt.Sprintf("This is %s %s. Press C-d to exit.\r\n", GoshName, GoshVersion)))

	for {
		n, err = MyGosh.term.Read(buf)
		if err != nil {
			// Consider MyGosh.term broken:
			os.Stdout.WriteString("Error reading input\n")
		} else {
			if buf[0] == 0x04 { // C-d to quit
				retcode = 0
				return
			} else {
				keys = utf8.ParseUTF8(buf[:n])
				for _, key = range keys {
					MyGosh.p.OnKey(key)
				}
			}
		}
	}

}
