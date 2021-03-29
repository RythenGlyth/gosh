package parser

import "gosh/src/shared"

type ExecutionStatement struct {
	method    ConstIdentifierValueStatement
	arguments []ValueStatement
}

func (es *ExecutionStatement) Exec(gosh shared.IGosh) {

}
