package parser

import "gosh/src/shared"

type IfStatement struct {
	conditionStatement ConditionStatement
	statementYes       ExecutableStatement
	statementNo        ExecutableStatement
}

func (is *IfStatement) Exec(gosh shared.IGosh) ValueStatement {
	if is.conditionStatement.EvalCondition() {
		if is.statementYes != nil {
			return is.statementYes.Exec(gosh)
		}
	} else {
		if is.statementNo != nil {
			return is.statementNo.Exec(gosh)
		}
	}
	return &ConstVoidValueStatement{}
}
