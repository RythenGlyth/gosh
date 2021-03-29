package parser

import "gosh/src/shared"

type LoopStatement struct {
}

func (ls *LoopStatement) Exec(gosh shared.IGosh) ValueStatement {
	return &ConstVoidValueStatement{}
}
