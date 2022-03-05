package bkill

type ErrUnknownSignal struct {
	Signal string
}

func (e *ErrUnknownSignal) Error() string {
	return "unknown signal: " + e.Signal
}
