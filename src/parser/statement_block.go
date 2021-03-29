package parser

import "gosh/src/shared"

type BlockStatement struct {
	statements []ExecutableStatement
}

func (bs *BlockStatement) Exec(gosh shared.IGosh) {
	for _, es := range bs.statements {
		es.Exec(gosh)
	}
}
