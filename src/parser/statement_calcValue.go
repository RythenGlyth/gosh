package parser

import (
	"math"
	"strings"
)

type CalcValueStatement struct {
	left    ValueStatement
	right   ValueStatement
	operand OperandType
}

func (cvs *CalcValueStatement) ValueType() ValueType {
	if cvs.left.ValueType() == VTString || cvs.right.ValueType() == VTString {
		return VTString
	}
	return VTNumber
}

func (cvs *CalcValueStatement) GetValue() (interface{}, ParseError) {
	if cvs.left.ValueType() == VTString || cvs.right.ValueType() == VTString {
		switch cvs.operand {
		case OTPlus:
			return cvs.left.String() + cvs.right.String(), nil
		case OTMinus:
			return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
		case OTMultiply:
			if cvs.left.ValueType() == VTNumber {
				return strings.Repeat(cvs.right.GetValue().(string), int(cvs.left.GetValue().(float64))), nil
			} else if cvs.right.ValueType() == VTNumber {
				return strings.Repeat(cvs.left.GetValue().(string), int(cvs.right.GetValue().(float64))), nil
			} else {
				return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
			}
		case OTDivide:
			return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
		}
	} else if cvs.left.ValueType() == VTNumber && cvs.right.ValueType() == VTNumber {
		l, _ := cvs.left.GetValue()
		lF := l.(float64)
		r, _ := cvs.right.GetValue()
		rF := r.(float64)
		switch cvs.operand {
		case OTPlus:
			return lF + rF, nil
		case OTMinus:
			return lF - rF, nil
		case OTMultiply:
			return lF * rF, nil
		case OTDivide:
			return lF / rF, nil
		case OTModulo:
			return math.Mod(lF, rF), nil
		}
	} else {
		return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
	}
}

type OperandType uint8

const (
	OTPlus OperandType = iota
	OTMinus
	OTMultiply
	OTDivide
	OTModulo
)

func (ot OperandType) String() string {
	switch ot {
	case OTPlus:
		return "+"
	case OTMinus:
		return "-"
	case OTMultiply:
		return "*"
	case OTDivide:
		return "/"
	case OTModulo:
		return "%"
	default:
		return "invalid"
	}
}
