package gosh

import (
	"fmt"
	"gosh/src/debug"
	"gosh/src/shared"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/scrouthtv/termios"
)

var keyCd termios.Key = termios.Key{Type: termios.KeyLetter, Mod: termios.ModCtrl, Value: 'd'}

// Gosh type collects all modules of a gosh shell.
type Gosh struct {
	term   termios.Terminal
	prompt *Prompt
	ready  bool
	debug  *debug.Client
	plugin *Handler
}

// InitError is an error that occurred during initialization of the gosh.
type InitError struct {
	err error
}

func (e *InitError) Unwrap() error {
	return e.err
}

func (e *InitError) Error() string {
	return "error initializing gosh: " + e.err.Error()
}

// NewGosh creates a new, empty gosh but does not start it yet.
func NewGosh() *Gosh {
	return &Gosh{nil, nil, false, nil, nil}
}

// SetDebugClient attaches the specified debugging client to
// the gosh instance.
func (g *Gosh) SetDebugClient(c *debug.Client) {
	g.debug = c
}

// LoadPlugin loads the plugin specified by the path.
// It returns an error if the file could not be read.
func (g *Gosh) LoadPlugin(s string) error {
	return g.plugin.Load(s)
}

// SendKey sends the specified key to all loaded plugins.
func (g *Gosh) SendKey(k *termios.Key) bool {
	return g.plugin.OnKey(k)
}

// DebugMessage sends a message with the specified module identifier
// and contents to the debug server.
func (g *Gosh) DebugMessage(k int, msg string) {
	g.debug.SendMessage(k, msg)
}

// Init prepares all sub-functionality of this gosh instance.
// If the returned error is not nil, this gosh must not be used until
// the issue is fixed.
func (g *Gosh) Init() error {
	var err error
	var term termios.Terminal

	term, err = termios.Open()
	if err != nil {
		return &InitError{err}
	}

	g.term = term

	g.prompt = NewPrompt(g)

	g.ready = true

	g.plugin = NewHandler(g)

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

// Interactive sets this gosh's mode to interactive.
// User input is read from the underlying terminal
// and commands are executed in the current namespace.
// A suitable exit code and any error is returned in the end.
func (g *Gosh) Interactive() (int, error) {
	var err error

	if !g.ready {
		err = g.Init()
		if err != nil {
			os.Stdout.WriteString("Error occurred during intialization:")
			os.Stdout.WriteString(err.Error())
			os.Stdout.WriteString("\n")

			return 1, err
		}
		defer g.Close()
	}

	var in []termios.Key
	var k termios.Key

	g.DebugMessage(shared.ModMain, "Going interactive")
	g.term.Write([]byte(fmt.Sprintf("This is %s %s. Press C-d to exit.\r\n", shared.GoshName, shared.GoshVersion)))
	g.prompt.redraw()

	for {
		g.term.SetRaw(true)
		in, err = g.term.Read()
		g.term.SetRaw(false)

		if err != nil {
			// Consider g.term broken:
			os.Stdout.WriteString("Error reading input:\n")
			os.Stdout.WriteString(err.Error() + "\n")
		} else {
			for _, k = range in {
				if k.Equal(&keyCd) { // C-d to quit
					return 0, nil
				}

				if !g.plugin.OnKey(&k) {
					g.DebugMessage(shared.ModMain, "Skipping this key because a plugin skipped it")
					continue
				}

				g.prompt.OnKey(k)

				if !g.ready {
					// If the terminal has been closed because of that key,
					// return:
					return 0, nil
				}
			}
		}
	}
}

// Eval evaluates the specified statement in the current namespace.
func (g *Gosh) Eval(line string) {
	parts := strings.Split(line, " ")
	if len(parts) == 0 || parts[0] == "" {
		return
	}

	switch parts[0] {
	case "exit":
		g.WriteString("Goodbye.")
		g.Close()
	case "cd":
		var err error

		if len(parts) == 1 {
			home, _ := os.UserHomeDir() //nolint errcheck // This will get implemented later in conjunction with the parser
			err = g.changeWD(home)
		} else {
			err = g.changeWD(parts[1])
		}

		if err != nil {
			g.WriteString("Error: ")
			g.WriteString(err.Error())
			g.WriteString("\r\n")
		}
	case "gst":
		cmd := exec.Command("git", "status")

		inPipe, inErr := cmd.StdinPipe()
		outPipe, outErr := cmd.StdoutPipe()
		errPipe, errErr := cmd.StderrPipe()

		if inErr != nil || outErr != nil || errErr != nil {
			g.WriteString("Error running the command: \r\n")
			g.WriteString(inErr.Error() + "\r\n")
			g.WriteString(outErr.Error() + "\r\n")
			g.WriteString(errErr.Error() + "\r\n")
		}

		go io.Copy(inPipe, os.Stdin)
		go io.Copy(os.Stdout, outPipe)
		go io.Copy(os.Stdout, errPipe)

		cmd.Start()
		cmd.Wait()
	default:
		g.WriteString("Unknown command '")
		g.WriteString(line)
		g.WriteString("'")
		g.WriteString("\r\n")
	}
}

func (g *Gosh) changeWD(target string) error {
	return os.Chdir(target)
}

// GetWD returns a string representation of the current working directory.
// ~ is returned if any error occurred (e. g. the directory was deleted).
func (g *Gosh) GetWD() string {
	wd, err := os.Getwd()
	if err != nil {
		return "~"
	}

	return wd
}

func (g *Gosh) Write(p []byte) (int, error) {
	return g.term.Write(p)
}

// WriteString writes the specified string to the gosh's terminal.
func (g *Gosh) WriteString(s string) (int, error) {
	return g.term.WriteString(s)
}
