package lexer

import "fmt"

func (lex *Lexer) readNumber(radix float64) (float64, LexError) {
	var currentVal float64
loop:
	for {
		if lex.position+1 >= lex.length {
			break loop
		}
		if lex.buffer[lex.position+1] == '.' {
			lex.next()
			var decVal float64
		decLoop:
			for {
				thisVal := getNumberValue(lex.buffer[lex.position+1])
				if thisVal > 0 && thisVal < radix {
				} else {
					break decLoop
				}
			}
			currentVal += decVal
			break loop
		} else {
			currentVal *= radix
			thisVal := getNumberValue(lex.buffer[lex.position+1])
			fmt.Print("_" + string(lex.buffer[lex.position+1]) + "_" + fmt.Sprint(thisVal) + ": " + fmt.Sprint(thisVal) + "\n")
			if thisVal >= 0 && thisVal < radix {
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
