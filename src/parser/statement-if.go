package parser

import (
	"fmt"
	"gosh/src/lexer"
	"io"
	"strings"
)

type IfStmt struct {
	cond   EvalStmt
	do     EvalStmt
	doElse EvalStmt
}

// read If statement starting with "if"
func (p *Parser) readIfStmt() (*IfStmt, ParseError) {
	i := IfStmt{}

	if p.tokens[p.pos].TokenType != lexer.TtIf {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, "if"}
	}
	advErr := p.tryAdvancePos() // advance onto (
	if advErr != nil {
		return nil, advErr
	}

	if p.tokens[p.pos].TokenType != lexer.TtLParen {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, "("}
	}

	var err error
	i.cond, err = p.readEvalStmt()
	if err != nil {
		return nil, err
	}
	advErr = p.tryAdvancePos() //advance onto )
	if advErr != nil {
		return nil, advErr
	}

	if p.tokens[p.pos].TokenType != lexer.TtRParen {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, ")"}
	}

	i.do, err = p.readEvalStmt()
	if err != nil {
		return nil, err
	}

	if p.tokens[p.pos].TokenType == lexer.TtElse {
		i.doElse, err = p.readEvalStmt()
		if err != nil {
			return nil, err
		}
	}

	return &i, nil
}

func (b *IfStmt) Eval() (v Value, err RuntimeError) {
	cV, cErr := b.cond.Eval()
	if cErr != nil {
		return nil, cErr
	}
	if cV.Bool() {
		return b.do.Eval()
	} else {
		if b.doElse != nil {
			return b.doElse.Eval()
		}
	}

	return &NilValue{}, nil
}

func (f *IfStmt) Debug(out io.Writer, indent int, symbol string) {
	fmt.Fprint(out, strings.Repeat(" ", indent))
	fmt.Fprintf(out, "%sIf:\n", symbol)
	f.cond.Debug(out, indent+1, "?")
	f.do.Debug(out, indent+1, "+")
	if f.doElse != nil {
		f.doElse.Debug(out, indent+1, ":")
	}
}
