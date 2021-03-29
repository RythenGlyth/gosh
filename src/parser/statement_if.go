package parser

import "gosh/src/shared"

type IfStatement struct {
	conditionStatement ConditionStatement
	statementYes       ExecutableStatement
	statementNo        ExecutableStatement
}

func (is *IfStatement) Exec(gosh shared.IGosh) {
	if is.conditionStatement.EvalCondition() {
		if is.statementYes != nil {
			is.statementYes.Exec(gosh)
		}
	} else {
		if is.statementNo != nil {
			is.statementNo.Exec(gosh)
		}
	}
}
