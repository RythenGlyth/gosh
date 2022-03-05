package main

import (
	"fmt"
	"gosh/src/builtins/bkill"
	"os"
	"strings"
)

func main() {
	g := NewPrintingGosh()
	k := bkill.Kill{}
	line := strings.Join(os.Args, " ")

	err := k.Eval(g, line)

	if err != nil {
		fmt.Println("uncaught error:", err)
	}
}
