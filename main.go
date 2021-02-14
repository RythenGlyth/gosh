package main

import (
	"io"
	"os"
)

func main() {
	var out io.WriteCloser
	var err error
	out, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0) //open our console (https://unix.stackexchange.com/questions/60641/linux-difference-between-dev-console-dev-tty-and-dev-tty0/60646#60646)
	if err != nil {
		os.Stdout.Write([]byte("Could not open /dev/tty for writing"))
	}
	defer out.Close()
	var in io.ReadCloser
	in, err = os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	if err != nil {
		os.Stdout.Write([]byte("Could not open /dev/tty for reading"))
	}
	defer in.Close()

	var inBuf []byte = make([]byte, 1024)
	var byteCount int
	for {
		byteCount, err = in.Read(inBuf)
		if err != nil {
			out.Write([]byte(err.Error()))
		} else {
			out.Write(inBuf[0:byteCount])
		}
	}

}
