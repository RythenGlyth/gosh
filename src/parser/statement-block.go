package parser

import (
	"fmt"
	"gosh/src/lexer"
	"io"
	"strings"
)

type BlockStmt struct {
	stmts []EvalStmt
}

// read block statement containing {
func (p *Parser) readBlock() (*BlockStmt, ParseError) {
	b := BlockStmt{}

	if p.pos+1 >= len(p.tokens) { //stop scanning if next is out of bounds
		return &b, nil
	}

	for p.tokens[p.pos+1].TokenType != lexer.TtRBrace {
		s, err := p.readEvalStmt()
		if err != nil {
			return nil, err
		}

		b.stmts = append(b.stmts, s)

		if p.pos+1 >= len(p.tokens) { //stop scanning if next is out of bounds
			return &b, nil
		}

		if p.tokens[p.pos+1].TokenType == lexer.TtSemicolon { //skip semicolons
			advErr := p.tryAdvancePos()
			if advErr != nil {
				return nil, advErr
			}
		}
		if p.pos+1 >= len(p.tokens) { //stop scanning if next is out of bounds
			return &b, nil
		}
	}

	advErr := p.tryAdvancePos() //advance onto }
	if advErr != nil {
		return nil, advErr
	}

	return &b, nil
}

func (b *BlockStmt) Eval() (v Value, err RuntimeError) {
	for _, s := range b.stmts {
		v, err = s.Eval()
		if err != nil {
			return nil, err
		}
	}

	return
}

func (b *BlockStmt) Debug(out io.Writer, indent int, symbol string) {
	fmt.Fprint(out, strings.Repeat(" ", indent))
	fmt.Fprintf(out, "%sBlock:\n", symbol)
	for _, s := range b.stmts {
		s.Debug(out, indent+1, "-")
	}
}
