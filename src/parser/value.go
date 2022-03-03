package parser

// Value of an EvalStmt is determined at runtime.
type Value interface {
	Bool() bool
	String() string
}

type NilValue struct {
}

func (nv *NilValue) Bool() bool {
	return false
}

func (nv *NilValue) String() string {
	return "null"
}
