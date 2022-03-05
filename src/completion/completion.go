package completion

import "gosh/src/shared"

type CompletionManager struct {
	paramCompleters map[string][]shared.IParamCompleter
	gosh            shared.IGosh
}

func NewCompletionManager(gosh shared.IGosh) *CompletionManager {
	return &CompletionManager{
		make(map[string][]shared.IParamCompleter),
		gosh,
	}
}

func (m *CompletionManager) RegisterParamCompleter(program string, c shared.IParamCompleter) {
	m.paramCompleters[program] = append(m.paramCompleters[program], c)
}

func (m *CompletionManager) CompleteLine(line string, position int) []shared.CompletionResult {
	results := make([]shared.CompletionResult, 0)

	// TODO

	return results
}

// CompleteParams attemps to complete the parameters of a program using a completer that has
// been set. If no completer could be found for this program, it returns nil.
func (m *CompletionManager) CompleteParams(program string, line string, position int) []shared.CompletionResult {
	results := make([]shared.CompletionResult, 16)

	for p, cs := range m.paramCompleters {
		if p == program {
			for _, c := range cs {
				results = append(results, c.CompleteParams(m.gosh, line, position)...)
			}
		}
	}

	return results
}
