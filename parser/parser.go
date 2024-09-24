package parser

import (
	"errors"
	"github.com/nam9nine/interpreter/ast"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
	"log"
)

func ParseProgram(l *lexer.Lexer) []ast.Statement {
	statements := []ast.Statement{}

	for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
		st, err := ParseStatement(t, l)
		if st == nil {
			log.Fatal("pointer nil error", err)
			return nil
		}
		statements = append(statements, st)
	}
	return statements
}

func ParseStatement(tok token.Token, l *lexer.Lexer) (ast.Statement, error) {
	tp := tok.Type
	switch tp {
	case token.LET:
		letSt, err := ParseLetStatement(tok, l)
		if err != nil {
			return nil, err
		}
		return letSt, nil
	case token.RETURN:
		ParseReturnStatement()
	default:
		return nil, nil
	}
	return nil, nil
}

func ParseLetStatement(tok token.Token, l *lexer.Lexer) (*ast.LetStatement, error) {
	letState := new(ast.LetStatement)
	identi := new(ast.Identifier)
	letState.Token = tok
	if t := l.NextToken(); t.Type == token.IDENT {
		identi.Value = t.Literal
		identi.Token = t
		letState.Name = identi
	} else {
		return nil, errors.New("식별자 없음")
	}

	if t := l.NextToken(); t.Type == token.ASSIGN {
	} else {
		return nil, errors.New("할당 연산자 없음")
	}

	if t := l.NextToken(); t.Type == token.INT {

	} else {
		return nil, errors.New("정수 없음")
	}

	if t := l.NextToken(); t.Type == token.SEMICOLON {
		// l.ch = ';' readchar = len(input)
	} else {
		return nil, errors.New("세미콜론 없음")
	}
	return letState, nil
}

func ParseReturnStatement() {
}
