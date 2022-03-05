package shared

// ICompletionManager handles completers and makes their functionality available.
type ICompletionManager interface {

	// RegisterParamCompleter registers a completer for usage with one specific program.
	RegisterParamCompleter(string, IParamCompleter)

	// CompleteLine completes something in the current line given the line
	// and position in the line.
	CompleteLine(string, int) []CompletionResult

	// CompleteParams completes the parameters of a program given the program name,
	// current command line and position in the line.
	CompleteParams(string, string, int) []CompletionResult
}

// IParamCompleter is used for completing the parameters of one specific program.
type IParamCompleter interface {
	CompleteParams(IGosh, string, int) []CompletionResult
}

type CompletionResult struct {
	Value   string
	Comment string
	Score   int
}
