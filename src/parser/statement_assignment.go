package parser

import (
	"gosh/src/lexer"
	"gosh/src/shared"
)

type AssignmentStatement struct {
	variableToken lexer.Token
}

func (as *AssignmentStatement) Exec(gosh shared.IGosh) {

}
