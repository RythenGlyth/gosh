package parser

import (
	"fmt"
	"io"
	"strings"
)

type ConstStmt struct {
	val Value
}

func (c *ConstStmt) Eval() (v Value, err RuntimeError) {
	return c.val, nil
}

func (f *ConstStmt) Debug(out io.Writer, indent int, symbol string) {
	fmt.Fprint(out, strings.Repeat(" ", indent))
	fmt.Fprintf(out, "%sConst: %s\n", symbol, f.val.String())
}
