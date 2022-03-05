package builtins

import (
	"gosh/src/shared"
	"strings"
)

type Which struct{}

func (b *Which) Match(line string) bool {
	return strings.HasPrefix(line, "which")
}

func (b *Which) Eval(g shared.IGosh, line string) error {
	// TODO

	return nil
}
