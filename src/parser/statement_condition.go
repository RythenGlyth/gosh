package parser

type ConditionStatement interface {
	EvalCondition() bool
}

type CompositeConditionStatement struct {
	a             ConditionStatement
	b             ConditionStatement
	conditionType CompositeConditionType
}

func (ccs *CompositeConditionStatement) EvalCondition() bool {
	switch ccs.conditionType {
	case CTAnd:
		return ccs.a.EvalCondition() && ccs.b.EvalCondition()
	case CTOr:
		return ccs.a.EvalCondition() || ccs.b.EvalCondition()
	}
	return false
}

type CompareConditionStatement struct {
	a             ValueStatement
	b             ValueStatement
	conditionType CompareConditionType
}

func (ccs *CompareConditionStatement) EvalCondition() bool {
	if ccs.a.ValueType() != ccs.b.ValueType() {
		return false
	}
	switch ccs.a.ValueType() {
	case VTNumber:
		switch ccs.conditionType {
		case CTEquals:
			return ccs.a.GetValue().(float64) == ccs.b.GetValue().(float64)
		case CTGreater:
			return ccs.a.GetValue().(float64) > ccs.b.GetValue().(float64)
		case CTGreaterEquals:
			return ccs.a.GetValue().(float64) >= ccs.b.GetValue().(float64)
		case CTLess:
			return ccs.a.GetValue().(float64) < ccs.b.GetValue().(float64)
		case CTLessEquals:
			return ccs.a.GetValue().(float64) <= ccs.b.GetValue().(float64)
		case CTNotEquals:
			return ccs.a.GetValue().(float64) != ccs.b.GetValue().(float64)
		}
	case VTString:
		switch ccs.conditionType {
		case CTEquals:
			return ccs.a.GetValue().(string) == ccs.b.GetValue().(string)
		case CTGreater:
			return ccs.a.GetValue().(string) > ccs.b.GetValue().(string)
		case CTGreaterEquals:
			return ccs.a.GetValue().(string) >= ccs.b.GetValue().(string)
		case CTLess:
			return ccs.a.GetValue().(string) < ccs.b.GetValue().(string)
		case CTLessEquals:
			return ccs.a.GetValue().(string) <= ccs.b.GetValue().(string)
		case CTNotEquals:
			return ccs.a.GetValue().(string) != ccs.b.GetValue().(string)
		}
	}
	return false
}

type ConstantConditionStatement struct {
	val bool
}

func (ccs *ConstantConditionStatement) EvalCondition() bool {
	return ccs.val
}

type NotConstantConditionStatement struct {
	a ConditionStatement
}

func (nccs *NotConstantConditionStatement) EvalCondition() bool {
	return !nccs.EvalCondition()
}

type CompositeConditionType uint8

const (
	CTAnd CompositeConditionType = iota
	CTOr
)

type CompareConditionType uint8

const (
	CTEquals CompareConditionType = iota
	CTNotEquals
	CTLessEquals
	CTGreaterEquals
	CTLess
	CTGreater
)
