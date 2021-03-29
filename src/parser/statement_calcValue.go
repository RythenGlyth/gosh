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
				l, lErr := cvs.left.GetValue()
				if lErr != nil {
					return nil, lErr
				}
				lF := l.(float64)
				r, rErr := cvs.right.GetValue()
				if rErr != nil {
					return nil, rErr
				}
				rS := r.(string)
				return strings.Repeat(rS, int(lF)), nil
			} else if cvs.right.ValueType() == VTNumber {
				l, lErr := cvs.left.GetValue()
				if lErr != nil {
					return nil, lErr
				}
				lS := l.(string)
				r, rErr := cvs.right.GetValue()
				if rErr != nil {
					return nil, rErr
				}
				rF := r.(float64)
				return strings.Repeat(lS, int(rF)), nil
			} else {
				return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
			}
		case OTDivide:
			return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
		}
	} else if cvs.left.ValueType() == VTNumber && cvs.right.ValueType() == VTNumber {
		l, lErr := cvs.left.GetValue()
		if lErr != nil {
			return nil, lErr
		}
		lF := l.(float64)
		r, rErr := cvs.right.GetValue()
		if rErr != nil {
			return nil, rErr
		}
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
	}
	return nil, &UnsupportedOperationError{cvs.left.ValueType(), cvs.right.ValueType(), cvs.operand}
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
