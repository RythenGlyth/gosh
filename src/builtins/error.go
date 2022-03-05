package builtins

type ErrMissingArg struct {
	Command string
	Missing string
}

func (e *ErrMissingArg) Error() string {
	return e.Command + ": missing argument " + e.Missing
}
