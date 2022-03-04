package parser

import (
	"gosh/src/lexer"
	"io"
)

type AssignStmt struct {
	variableToken lexer.Token
	value         EvalStmt
}

func (b *AssignStmt) Eval() (v Value, err RuntimeError) {
	// TODO
	return b.value.Eval()
}

func (b *AssignStmt) Debug(out io.Writer, indent int, symbol string) {
	// fmt.Fprint(out, strings.Repeat(" ", indent))
	// fmt.Fprintf(out, "%sBlock:\n", symbol)
	// for _, s := range b.stmts {
	// 	s.Debug(out, indent+1, "-")
	// }
}
