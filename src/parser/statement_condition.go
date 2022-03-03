package parser

import "strconv"

type ConditionStatement interface {
	EvalCondition() (bool, ParseError)
}

type CompositeConditionStatement struct {
	a             ConditionStatement
	b             ConditionStatement
	conditionType CompositeConditionType
}

func (ccs *CompositeConditionStatement) EvalCondition() (bool, ParseError) {
	switch ccs.conditionType {
	case CTAnd:
		boolA, errA := ccs.a.EvalCondition()
		if errA != nil {
			return false, errA
		}
		if boolA {
			boolB, errB := ccs.b.EvalCondition()
			if errB != nil {
				return false, errB
			}
			return boolB, nil
		}
		return false, nil
	case CTOr:
		boolA, errA := ccs.a.EvalCondition()
		if errA != nil {
			return false, errA
		}
		if boolA {
			return true, nil
		} else {
			boolB, errB := ccs.b.EvalCondition()
			if errB != nil {
				return false, errB
			}
			return boolB, nil
		}
	}
	return false, nil
}

func (ccs *CompositeConditionStatement) ValueType() ValueType {
	return VTBoolean
}

func (ccs *CompositeConditionStatement) GetValue() (interface{}, ParseError) {
	val, err := ccs.EvalCondition()
	return val, err
}

func (ccs *CompositeConditionStatement) String() string {
	val, err := ccs.EvalCondition()
	if err != nil {
		return "booleanError"
	} else {
		return strconv.FormatBool(val)
	}
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

func (ccs *CompareConditionStatement) ValueType() ValueType {
	return VTBoolean
}

func (ccs *CompareConditionStatement) GetValue() (interface{}, ParseError) {
	val, err := ccs.EvalCondition()
	return val, err
}

func (ccs *CompareConditionStatement) String() string {
	val, err := ccs.EvalCondition()
	if err != nil {
		return "booleanError"
	} else {
		return strconv.FormatBool(val)
	}
}

type ConstantConditionStatement struct {
	val bool
}

func (ccs *ConstantConditionStatement) EvalCondition() (bool, ParseError) {
	return ccs.val, nil
}

func (ccs *ConstantConditionStatement) ValueType() ValueType {
	return VTBoolean
}

func (ccs *ConstantConditionStatement) GetValue() (interface{}, ParseError) {
	val, err := ccs.EvalCondition()
	return val, err
}

func (ccs *ConstantConditionStatement) String() string {
	val, err := ccs.EvalCondition()
	if err != nil {
		return "booleanError"
	} else {
		return strconv.FormatBool(val)
	}
}

type NotConstantConditionStatement struct {
	a ConditionStatement
}

func (nccs *NotConstantConditionStatement) EvalCondition() (bool, ParseError) {
	val, err := nccs.EvalCondition()
	if err != nil {
		return false, err
	}
	return !val, nil
}

func (ccs *NotConstantConditionStatement) ValueType() ValueType {
	return VTBoolean
}

func (ccs *NotConstantConditionStatement) GetValue() (interface{}, ParseError) {
	val, err := ccs.EvalCondition()
	return val, err
}

func (ccs *NotConstantConditionStatement) String() string {
	val, err := ccs.EvalCondition()
	if err != nil {
		return "booleanError"
	} else {
		return strconv.FormatBool(val)
	}
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
