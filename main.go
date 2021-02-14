package main

import (
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/google/goterm/term"
)

func main() {
	var err error
	var backupTerm term.Termios
	backupTerm, err = term.Attr(os.Stdin)
	if err != nil {
		os.Stdout.Write([]byte("Could not copy Stdin attributes into TTY: " + err.Error()))
		os.Exit(5)
	}
	myTerm := backupTerm
	myTerm.Raw()
	myTerm.Set(os.Stdin)

	defer backupTerm.Set(os.Stdin)

	sig := make(chan os.Signal, 2)

	os.Stdout.WriteString("STARTED")

	signal.Notify(sig, unix.SIGWINCH, unix.SIGCLD)

	os.Stdout.WriteString("STARTED2")

	myTerm.Winsz(os.Stdin)

	for {
		var buf = make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			os.Stdout.WriteString("E")
		} else {
			os.Stdout.Write([]byte(fmt.Sprintf("W: %s\n\r", buf[0:n])))
		}

	}

}

/*func writer() {
	var buf = make([]byte, bufSz)
}*/
