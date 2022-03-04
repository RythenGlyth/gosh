package builtins

import (
	"gosh/src/shared"
	"strings"
)

type BuiltinAlias struct{}

func (b *BuiltinAlias) Match(line string) bool {
	return strings.HasPrefix(line, "alias")
}

func (b *BuiltinAlias) Eval(g shared.IGosh, line string) error {
	words := strings.SplitN(line, " ", 4)
	if len(words) == 0 {
		return &ErrAliasBadUsage{"matcher"} // TODO print all aliases instead
	}

	global := words[1] == "-g" || words[1] == "--global"
	glen := 0
	if global {
		glen = len(words[1]) + 1
		copy(words[1:], words[2:]) // remove global flag from params
		words = words[:len(words)-1]
	}

	split := strings.IndexRune(words[1], '=')

	var m, e string

	if split == -1 {
		// attempt to split at space instead
		if strings.Count(line, " ") > 3 || (!global && strings.Count(line, " ") > 2) {
			return &ErrAliasBadUsage{"expansion"} // TODO print usage
		}

		m = words[1]
		e = strings.Join(words[2:], " ")
	} else {
		m = words[1][:split]
		split += len(words[0]) + 1 + glen
		e = line[split+1:]
	}

	println("aliasing", m, "to", e)
	if global {
		g.GetAliasManager().RegisterGlobalAlias(m, e)
	} else {
		g.GetAliasManager().RegisterSimpleAlias(m, e)
	}

	return nil
}

type ErrAliasBadUsage struct {
	Missing string
}

func (e *ErrAliasBadUsage) Error() string {
	return "missing " + e.Missing
}
