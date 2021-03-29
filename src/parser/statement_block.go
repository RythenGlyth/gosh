package parser

import "gosh/src/shared"

type BlockStatement struct {
	statements []ExecutableStatement
}

func (bs *BlockStatement) Exec(gosh shared.IGosh) ValueStatement {
	for _, es := range bs.statements {
		es.Exec(gosh)
	}
	return &ConstVoidValueStatement{} //TODO Real value
}
