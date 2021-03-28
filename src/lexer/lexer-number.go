package lexer

import (
	"math"
)

func (lex *Lexer) readNumber(radix float64) (float64, LexError) {
	var currentVal float64
	thisVal := getNumberValue(lex.buffer[lex.position])
	if thisVal >= 0 && thisVal < radix {
		currentVal *= radix
		currentVal += thisVal
	} else {
		return 0, nil
	}
loop:
	for lex.position+1 < lex.length {
		if lex.buffer[lex.position+1] == '.' {
			lex.next()
			//var decVal float64
			var i float64 = 1
		decLoop:
			for lex.position+1 < lex.length {
				thisVal := getNumberValue(lex.buffer[lex.position+1])
				if thisVal >= 0 && thisVal < radix {
					//decVal *= radix
					//decVal += thisVal
					currentVal += thisVal / math.Pow(radix, i)
					lex.next()
					i++
				} else {
					break decLoop
				}
			}
			//currentVal += decVal
			break loop
		} else {
			thisVal := getNumberValue(lex.buffer[lex.position+1])
			if thisVal >= 0 && thisVal < radix {
				currentVal *= radix
				currentVal += thisVal
				lex.next()
			} else {
				break loop
			}
		}
	}

	return currentVal, nil
}

func getNumberValue(char rune) float64 {
	if char >= '0' && char <= '9' {
		return (float64)(char - '0')
	}

	if char >= 'A' && char <= 'Z' {
		return (float64)(char - 'A' + 10)
	}

	if char >= 'a' && char <= 'z' {
		return (float64)(char - 'a' + 10)
	}

	return -1
}
