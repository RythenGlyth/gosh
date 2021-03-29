package interpreter

import "io/ioutil"

type Interpreter struct {
	contents []rune
}

func InterpreterByFile(filename string) (*Interpreter, error) {
	buf, err := ioutil.ReadFile("../test/missingQuote.gosh")
	if err != nil {
		return nil, &ErrFileRead{err}
	}

	contents := []rune(string(buf))
	return &Interpreter{contents}, nil
}

func InterpreterByString(content string) *Interpreter {
	return &Interpreter{[]rune(content)}
}

func (i *Interpreter) Run() [Statements]error {
	//
}
