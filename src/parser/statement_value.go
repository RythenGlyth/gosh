package parser

import (
	"gosh/src/lexer"
	"strconv"
	"strings"
)

type ValueStatement interface {
	ValueType() ValueType
	GetValue() (interface{}, ParseError)
	String() string
}

type ConstNumberValueStatement struct {
	number float64
}

func (cnvs *ConstNumberValueStatement) ValueType() ValueType {
	return VTNumber
}

func (cnvs *ConstNumberValueStatement) GetValue() (interface{}, ParseError) {
	return cnvs.number, nil
}

func (cnvs *ConstNumberValueStatement) String() string {
	return strconv.FormatFloat(cnvs.number, 'G', -1, 64)
}

type ConstVoidValueStatement struct {
}

func (cvvs *ConstVoidValueStatement) ValueType() ValueType {
	return VTVoid
}

func (cvvs *ConstVoidValueStatement) GetValue() (interface{}, ParseError) {
	return nil, nil
}

func (cvvs *ConstVoidValueStatement) String() string {
	return "null"
}

type ConstStringValueStatement struct {
	text string
}

func (csvs *ConstStringValueStatement) ValueType() ValueType {
	return VTString
}

func (csvs *ConstStringValueStatement) GetValue() (interface{}, ParseError) {
	return csvs.text, nil
}

func (csvs *ConstStringValueStatement) String() string {
	return csvs.text
}

type ConstArrayValueStatement struct {
	array []ValueStatement
}

func (cavs *ConstArrayValueStatement) ValueType() ValueType {
	return VTArray
}

func (cavs *ConstArrayValueStatement) GetValue() (interface{}, ParseError) {
	return cavs.array, nil
}

func (cavs *ConstArrayValueStatement) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	length := len(cavs.array)
	for idx, el := range cavs.array {
		builder.WriteString(el.String())
		if idx+1 < length {
			builder.WriteRune(',')
			builder.WriteRune(' ')
		}
	}
	builder.WriteRune(']')
	return builder.String()
}

type ConstMapValueStatement struct {
	theMap map[ValueStatement]ValueStatement
}

func (cmvs *ConstMapValueStatement) ValueType() ValueType {
	return VTMap
}

func (cmvs *ConstMapValueStatement) GetValue() (interface{}, ParseError) {
	return cmvs.theMap, nil
}

type ConstIdentifierValueStatement struct {
	i string
}

func (civs *ConstIdentifierValueStatement) ValueType() ValueType {
	return VTIdentifer
}

func (civs *ConstIdentifierValueStatement) GetValue() (interface{}, ParseError) {
	return civs.i, nil
}

func (cmvs *ConstMapValueStatement) String() string {
	var builder strings.Builder
	builder.WriteRune('[')
	length := len(cmvs.theMap)
	idx := 1
	for key, val := range cmvs.theMap {
		builder.WriteString(key.String())
		builder.WriteRune(':')
		builder.WriteRune(' ')
		builder.WriteString(val.String())
		if idx < length {
			builder.WriteRune(',')
			builder.WriteRune(' ')
		}
		idx++
	}
	builder.WriteRune(']')
	return builder.String()
}

type VarValueStatement struct {
	variableToken lexer.Token
}

type ValueType uint8

const (
	VTNumber ValueType = iota
	VTString
	VTArray
	VTMap
	VTIdentifer
	VTBoolean
	VTVoid
)

func (vt ValueType) String() string {
	switch vt {
	case VTNumber:
		return "number"
	case VTArray:
		return "array"
	case VTMap:
		return "map"
	case VTString:
		return "string"
	case VTIdentifer:
		return "identifier"
	case VTBoolean:
		return "boolean"
	case VTVoid:
		return "void"
	default:
		return "invalid"
	}
}
