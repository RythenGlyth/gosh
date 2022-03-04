package builtins

import (
	"gosh/src/shared"
	"os"
	"strings"
)

type BuiltinCd struct{}

func (b *BuiltinCd) Match(line string) bool {
	return strings.HasPrefix(line, "cd")
}

func (b *BuiltinCd) Eval(g shared.IGosh, line string) error {
	parts := strings.Split(line, " ")

	if len(parts) == 1 { // no parameter
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		return g.ChangeWD(home)
	} else {
		return g.ChangeWD(parts[1])
	}
}
