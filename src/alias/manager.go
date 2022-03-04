package alias

import (
	"gosh/src/shared"
	"strings"
)

type Manager struct {
	simpleAliases []Alias
	globalAliases []Alias
}

func (m *Manager) RegisterSimpleAlias(match, expand string) {
	m.simpleAliases = append(m.simpleAliases, Alias{match, expand})
}

func (m *Manager) RegisterGlobalAlias(match, expand string) {
	m.globalAliases = append(m.globalAliases, Alias{match, expand})
}

func (m *Manager) Expand(line string) string {
	ok := true
	for ok { // expand as long as occurrences were found
		line, ok = m.expandSimply(line)
	}

	ok = true
	for ok {
		line, ok = m.expandGlobally(line)
	}

	return line
}

func (m *Manager) expandSimply(line string) (string, bool) {
	split := strings.Index(line, " ")
	if split == -1 {
		split = len(line)
	}

	word := line[:split]

	for _, a := range m.simpleAliases {
		if a.Match == word {
			return a.Expand + line[split:], true
		}
	}

	return line, false
}

func (m *Manager) expandGlobally(line string) (string, bool) {
	words := strings.Split(line, " ")

	for i, w := range words {
		for _, a := range m.globalAliases {
			if a.Match == w {
				pre := ""
				if i > 0 {
					pre = strings.Join(words[:i], " ") + " "
				}

				post := ""
				if i < len(words)-1 {
					post = " " + strings.Join(words[i+1:], " ")
				}

				return pre + a.Expand + post, true
			}
		}
	}

	return line, false
}

func (m *Manager) ListSimpleAliases() []shared.Alias {
	as := make([]shared.Alias, len(m.simpleAliases))
	for i := range m.simpleAliases {
		as[i] = &(m.simpleAliases[i])
	}

	return as
}

func (m *Manager) ListGlobalAliases() []shared.Alias {
	as := make([]shared.Alias, len(m.globalAliases))
	for i := range m.globalAliases {
		as[i] = &(m.globalAliases[i])
	}

	return as
}
