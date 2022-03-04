package builtins

import (
	"gosh/src/shared"
	"os"
	"strings"
)

type Cd struct{}

func (b *Cd) Match(line string) bool {
	return strings.HasPrefix(line, "cd")
}

func (b *Cd) Eval(g shared.IGosh, line string) error {
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
