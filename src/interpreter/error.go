package interpreter

type ErrFileRead struct {
	err error
}

func (err *ErrFileRead) Error() string {
	return "Error reading file: " + err.err.Error()
}

func (err *ErrFileRead) Unwrap() error {
	return err.err
}
