package builtins

import (
	"gosh/src/shared"
	"strings"
)

type BuiltinPwd struct{}

func (b *BuiltinPwd) Match(line string) bool {
	return strings.HasPrefix(line, "pwd")
}

func (b *BuiltinPwd) Eval(g shared.IGosh, line string) error {
	path, err := g.GetWD()
	if err != nil {
		return err
	}

	g.WriteString(path)
	return nil
}
