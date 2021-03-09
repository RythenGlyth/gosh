package debug

// LaunchError is returned if the debugger could not start.
type LaunchError struct {
	err error
}

func (e *LaunchError) Unwrap() error {
	return e.err
}

func (e *LaunchError) Error() string {
	return "error starting debugger: " + e.err.Error()
}

// WriteError is returned if an error occurred whilst writing debugging data.
type WriteError struct {
	err error
}

func (e *WriteError) Unwrap() error {
	return e.err
}

func (e *WriteError) Error() string {
	return "error writing debugging data: " + e.err.Error()
}
