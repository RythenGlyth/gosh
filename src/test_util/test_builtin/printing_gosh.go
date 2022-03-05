package main

import (
	"fmt"
	"gosh/src/test_util"
	"io"
	"os"
)

type PrintingGosh struct {
	test_util.GoshStub
	out io.Writer
}

func (g *PrintingGosh) Write(buf []byte) (n int, err error) {
	return g.out.Write(buf)
}

func (g *PrintingGosh) WriteString(s string) (n int, err error) {
	return fmt.Fprint(g.out, s)
}

func NewPrintingGosh() *PrintingGosh {
	return &PrintingGosh{
		test_util.GoshStub{T: nil},
		os.Stdout,
	}
}
