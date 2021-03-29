package parser

import "gosh/src/lexer"

type ValueStatement interface {
	ValueType() ValueType
	GetValue() interface{}
}

type ConstNumberValueStatement struct {
	number float64
}

func (cnvs *ConstNumberValueStatement) ValueType() ValueType {
	return VTNumber
}

func (cnvs *ConstNumberValueStatement) GetValue() interface{} {
	return cnvs.number
}

type ConstStringValueStatement struct {
	text string
}

func (csvs *ConstStringValueStatement) ValueType() ValueType {
	return VTString
}

func (csvs *ConstStringValueStatement) GetValue() interface{} {
	return csvs.text
}

type VarValueStatement struct {
	variableToken lexer.Token
}

type ValueType uint8

const (
	VTNumber ValueType = iota
	VTString
)
