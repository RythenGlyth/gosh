package parser

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

func (cvs *CalcValueStatement) GetValue() interface{} {
	if cvs.left.ValueType() == VTString || cvs.right.ValueType() == VTString {
		switch cvs.operand {

		}
	} else {
		switch cvs.operand {
		case OTPlus:
			return cvs.left.GetValue().(float64) + cvs.right.GetValue().(float64)
		case OTMinus:
			return cvs.left.GetValue().(float64) - cvs.right.GetValue().(float64)
		case OTMultiply:
			return cvs.left.GetValue().(float64) * cvs.right.GetValue().(float64)
		case OTDivide:
			return cvs.left.GetValue().(float64) / cvs.right.GetValue().(float64)
		case OTModulo: //TODO: throw exception if not int
			return int64(cvs.left.GetValue().(float64)) % int64(cvs.right.GetValue().(float64))
		}
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
