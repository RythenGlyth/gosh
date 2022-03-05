package main

import (
	"fmt"
	"gosh/src/builtins/bkill"
	"gosh/src/gosh"
	"os"
	"strings"
)

func main() {
	g := gosh.NewGosh()
	g.Init()
	defer g.Close()

	k := bkill.Kill{}
	line := strings.Join(os.Args, " ")

	err := k.Eval(g, line)

	if err != nil {
		fmt.Println("uncaught error:", err)
	}
}
