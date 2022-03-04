package builtins

import (
	"gosh/src/shared"
	"strings"
)

type Unset struct{}

func (b *Unset) Match(line string) bool {
	return strings.HasPrefix(line, "unset")
}

func (b *Unset) Eval(g shared.IGosh, line string) error {
	// TODO
	return nil
}
