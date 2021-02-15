package main

import (
	"encoding/hex"
	"os"
	"os/signal"

	"github.com/google/goterm/term"
	"golang.org/x/sys/unix"
)

func main() {
	retcode := 0

	defer func() { os.Exit(retcode) }()
	var err error
	var backupTerm term.Termios
	backupTerm, err = term.Attr(os.Stdin)
	if err != nil {
		os.Stdout.Write([]byte("Could not copy Stdin attributes into TTY: " + err.Error()))
		retcode = 5
		return
	}
	myTerm := backupTerm
	myTerm.Raw()
	myTerm.Set(os.Stdin)

	defer backupTerm.Set(os.Stdin)

	sig := make(chan os.Signal, 2)

	signal.Notify(sig, unix.SIGWINCH, unix.SIGCLD)

	myTerm.Winsz(os.Stdin)

	for {
		var buf = make([]byte, 1024)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			os.Stdout.WriteString("Error reading input")
		} else {
			//os.Stdout.Write([]byte(fmt.Sprintf("W: %s\n\r", hex.EncodeToString(buf[0:n]))))
			os.Stdout.Write(buf[0:n])
			if buf[0] == 0x03 { //ctrl+c to quit
				retcode = 0
				return
			}
		}
	}

}

func byteArrToHex(arr []byte) []byte {
	out := make([]byte, hex.EncodedLen(len(arr)))
	hex.Encode(out, arr)

	return out
}
