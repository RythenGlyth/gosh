package util

import (
	"fmt"
	"io"
	"strings"
)

func PrintTable(f io.Writer, cells []string, width int) {
	cw := 0 // (largest) cell width
	for _, cell := range cells {
		if len(cell) > cw {
			cw = len(cell)
		}
	}
	println("width", width)

	brk := width / (cw + 2) // two extra spaces after all cells
	var i int
	var cell string

	for i, cell = range cells {
		fmt.Fprint(f, cell)
		fmt.Fprint(f, strings.Repeat(" ", cw-len(cell)))

		if i%brk == brk-1 || i == len(cells)-1 {
			fmt.Fprint(f, "\n")
		} else {
			fmt.Fprint(f, "  ")
		}
	}
}
