package builtins

import (
	"gosh/src/shared"
	"strings"
)

type BuiltinExit struct{}

func (b *BuiltinExit) Match(line string) bool {
	return strings.HasPrefix("exit", line)
}

func (b *BuiltinExit) Eval(g shared.IGosh, line string) error {
	g.WriteString("Goodbye.")
	g.Close()

	return nil
}
