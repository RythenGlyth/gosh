package builtins

import (
	"gosh/src/shared"
	"strings"
)

type Exit struct{}

func (b *Exit) Match(line string) bool {
	return strings.HasPrefix("exit", line)
}

func (b *Exit) Eval(g shared.IGosh, line string) error {
	g.WriteString("Goodbye.")
	g.Close()

	return nil
}
