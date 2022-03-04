package parser

import (
	"fmt"
	"gosh/src/lexer"
	"io"
	"strings"
)

type ForStmt struct {
	doPre  EvalStmt
	cond   EvalStmt
	doPost EvalStmt
	body   EvalStmt
}

// read For statement starting with "for"
func (p *Parser) readForStmt() (*ForStmt, ParseError) {
	f := ForStmt{}

	if p.tokens[p.pos].TokenType != lexer.TtFor {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, "for"}
	}
	advErr := p.tryAdvancePos() //advance onto (
	if advErr != nil {
		return nil, advErr
	}

	if p.tokens[p.pos].TokenType != lexer.TtLParen {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, "("}
	}

	var err RuntimeError
	f.cond, err = p.readEvalStmt()
	if err != nil {
		return nil, err
	}

	advErr = p.tryAdvancePos() //advance onto ; or )
	if advErr != nil {
		return nil, advErr
	}
	if p.tokens[p.pos].TokenType == lexer.TtSemicolon {
		f.doPre = f.cond
		f.cond = nil

		f.cond, err = p.readEvalStmt()
		if err != nil {
			return nil, err
		}
		advErr = p.tryAdvancePos() // advance onto second ;
		if advErr != nil {
			return nil, advErr
		}

		if p.tokens[p.pos].TokenType != lexer.TtSemicolon {
			return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, ";"}
		}

		f.doPost, err = p.readEvalStmt()
		if err != nil {
			return nil, err
		}
		advErr = p.tryAdvancePos() //advance onto )
		if advErr != nil {
			return nil, advErr
		}
	}

	if p.tokens[p.pos].TokenType != lexer.TtRParen {
		return nil, &UnexpectedTokenError{p.tokens[p.pos].TokenType, ")"}
	}

	f.body, err = p.readEvalStmt() // read the body
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (f *ForStmt) Eval() (v Value, err RuntimeError) {
	if f.doPre != nil {
		_, err = f.doPre.Eval()
		if err != nil {
			return nil, err
		}
	}

theFor:
	for {
		val, err := f.cond.Eval()
		if err != nil {
			return nil, err
		}
		if !val.Bool() {
			break theFor
		}

		v, err = f.body.Eval()
		if err != nil {
			return nil, err
		}

		if f.doPost != nil {
			_, err = f.doPost.Eval()
			if err != nil {
				return nil, err
			}
		}
	}

	return
}

func (f *ForStmt) Debug(out io.Writer, indent int, symbol string) {
	fmt.Fprint(out, strings.Repeat(" ", indent))
	fmt.Fprintf(out, "%sFor:\n", symbol)
	f.cond.Debug(out, indent+1, "?")
	f.body.Debug(out, indent+1, "+")
	if f.doPre != nil {
		f.doPre.Debug(out, indent+1, "<")
	}
	if f.doPost != nil {
		f.doPost.Debug(out, indent+1, ">")
	}
}
