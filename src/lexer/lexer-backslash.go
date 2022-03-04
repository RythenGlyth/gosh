package lexer

func (lex *Lexer) nextB(wasBackslash bool) LexError {
	lex.wasBackslash = wasBackslash
	lex.nextPos()

	if lex.position >= 0 && lex.position < lex.length && lex.buffer[lex.position] == '\\' && !lex.wasBackslash {
		lex.wasBackslash = true

		if lex.position+1 < lex.length {
			return lex.nextB(true)
		}

		return &TrailingBackslashError{Position{lex.codeXPos, lex.codeYPos, lex.position, lex}}
	}

	if lex.position >= 0 && lex.position < lex.length {
		if lex.wasBackslash {
			switch lex.buffer[lex.position] {
			case 'b':
				lex.character = '\b'
			case 'f':
				lex.character = '\f'
			case 'n':
				lex.character = '\n'
			case 'r':
				lex.character = '\r'
			case 't':
				lex.character = '\t'
			case 'v':
				lex.character = '\v'
			case '\\':
				lex.character = '\\'
			case '\'':
				lex.character = '\''
			case ' ':
				lex.character = ' '
			case '"':
				lex.character = '"'
			case 'x', 'u', 'U':
				lex.nextPos()

				val, err := lex.readNumber(16)
				if err != nil {
					return err
				}

				lex.character = rune(val)
			default:
				return &InvalidEscapeCharacterError{
					Position{lex.codeXPos, lex.codeYPos, lex.position, lex},
					lex.buffer[lex.position],
				}
			}
		} else {
			lex.character = lex.buffer[lex.position]
		}
	}

	return nil
}
