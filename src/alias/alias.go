package alias

type Alias struct {
	Match  string
	Expand string
}

func (a *Alias) Format() string {
	return a.Match + " aliased to " + a.Expand
}
