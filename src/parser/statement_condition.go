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

func (ccs *CompareConditionStatement) EvalCondition() (bool, ParseError) {
	if ccs.a.ValueType() != ccs.b.ValueType() {
		return false, nil
	}
	switch ccs.a.ValueType() {
	case VTNumber:
		l, lErr := ccs.a.GetValue()
		if lErr != nil {
			return false, lErr
		}
		lF := l.(float64)
		r, rErr := ccs.b.GetValue()
		if rErr != nil {
			return false, rErr
		}
		rF := r.(float64)
		switch ccs.conditionType {
		case CTEquals:
			return lF == rF, nil
		case CTGreater:
			return lF > rF, nil
		case CTGreaterEquals:
			return lF >= rF, nil
		case CTLess:
			return lF < rF, nil
		case CTLessEquals:
			return lF <= rF, nil
		case CTNotEquals:
			return lF != rF, nil
		}
	case VTString:
		l, lErr := ccs.a.GetValue()
		if lErr != nil {
			return false, lErr
		}
		lS := l.(string)
		r, rErr := ccs.b.GetValue()
		if rErr != nil {
			return false, rErr
		}
		rS := r.(string)
		switch ccs.conditionType {
		case CTEquals:
			return lS == rS, nil
		case CTGreater:
			return lS > rS, nil
		case CTGreaterEquals:
			return lS >= rS, nil
		case CTLess:
			return lS < rS, nil
		case CTLessEquals:
			return lS <= rS, nil
		case CTNotEquals:
			return lS != rS, nil
		}
	}
	return false, &UnsupportedConditionError{ccs.a.ValueType(), ccs.b.ValueType(), &ccs.conditionType}
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

type ConditionType interface {
	conditionTypeceb2a0f0028f64cf89ef7e4b()
}

type CompositeConditionType uint8

const (
	CTAnd CompositeConditionType = iota
	CTOr
)

func (cct *CompositeConditionType) conditionTypeceb2a0f0028f64cf89ef7e4b() {

}

func (cct CompositeConditionType) String() string {
	switch cct {
	case CTAnd:
		return "&&"
	case CTOr:
		return "||"
	default:
		return "invalid"
	}
}

type CompareConditionType uint8

const (
	CTEquals CompareConditionType = iota
	CTNotEquals
	CTLessEquals
	CTGreaterEquals
	CTLess
	CTGreater
)

func (cct *CompareConditionType) conditionTypeceb2a0f0028f64cf89ef7e4b() {

}

func (cct CompareConditionType) String() string {
	switch cct {
	case CTEquals:
		return "=="
	case CTNotEquals:
		return "!="
	case CTLessEquals:
		return "<="
	case CTGreaterEquals:
		return ">="
	case CTLess:
		return "<"
	case CTGreater:
		return ">"
	default:
		return "invalid"
	}
}
