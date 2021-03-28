package lexer

import (
	"fmt"
	"math"
)

func (lex *Lexer) readNumber(radix float64) (float64, LexError) {
	var currentVal float64
loop:
	for lex.position < lex.length {
		if lex.character == '.' {
			lex.next()
			//var decVal float64
			var i float64 = 1
		decLoop:
			for lex.position < lex.length {
				thisVal := getNumberValue(lex.character)
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
			thisVal := getNumberValue(lex.character)
			fmt.Print("_" + string(lex.character) + "_" + fmt.Sprint(thisVal) + ": " + fmt.Sprint(thisVal) + "\n")
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
