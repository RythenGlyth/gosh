package lexer

// Lexer (Tokenizer)
type Lexer struct {
	//buffer containing the character of the code to analyze
	buffer []byte
	//character inden in buffer
	position int
	//length of the content in the buffer
	length int
	//if the last character was a backslash. used to handle escaped codes
	wasBackslash bool
}

// NewLexer creates a new Lexer
func NewLexer(source []byte, length int) *Lexer {
	return &Lexer{source, 0, length, false}
}

// Lex into tokens
func (lex *Lexer) Lex() []Token {
	var tokens []Token
	for lex.position < lex.length {
		tokens = append(tokens, lex.nextToken())
	}
	return tokens
}

func (lex *Lexer) nextToken() Token {
	var startPos int = lex.position
	var tokenType TokenType = ttEmpty
loop:
	for {
		if lex.position >= lex.length {
			break loop
		}
		switch lex.buffer[lex.position] {
		case ' ', '\t', '\f', '\n', '\r':
			lex.position++
			break
		case ';', ',', '.', '(', ')', '[', ']', '{', '}', '%', '*':
			tokenType = TokenType(string(lex.buffer[lex.position : lex.position+1]))
			break loop
		default:
			break loop
		}
	}
	lex.position++
	var endpos int = lex.position
	return Token{tokenType, startPos, endpos}
}
