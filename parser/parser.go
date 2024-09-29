package parser

import (
	"errors"
	"github.com/nam9nine/interpreter/ast"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
)

type Parser struct {
	lex *lexer.Lexer
	err []error
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		lex: l,
		err: []error{},
	}
}

func (p *Parser) ParseProgram() []ast.Statement {
	statements := []ast.Statement{}

	for t := p.lex.NextToken(); t.Type != token.EOF; t = p.lex.NextToken() {
		st := p.ParseStatement(t)
		statements = append(statements, st)
	}
	return statements
}

func (p *Parser) ParseStatement(tok token.Token) ast.Statement {
	var statement ast.Statement

	tp := tok.Type
	switch tp {
	case token.LET:
		statement = p.ParseLetStatement(tok)
	case token.RETURN:
		statement = ParseReturnStatement()
	default:
		return nil
	}
	return statement
}

func (p *Parser) ParseLetStatement(tok token.Token) *ast.LetStatement {
	letState := new(ast.LetStatement)
	identi := new(ast.Identifier)

	letState.Token = tok

	if t := p.lex.NextToken(); t.Type == token.IDENT {
		identi.Value = t.Literal
		identi.Token = t
		letState.Name = identi
	} else {
		p.err = append(p.err, errors.New("식별자 없음"))
	}

	if t := p.lex.NextToken(); t.Type == token.ASSIGN {
	} else {
		p.err = append(p.err, errors.New("할당 연산자 없음"))
	}

	if t := p.lex.NextToken(); t.Type == token.INT {

	} else {
		p.err = append(p.err, errors.New("정수 없음"))
	}

	if t := p.lex.NextToken(); t.Type == token.SEMICOLON {
	} else {
		p.err = append(p.err, errors.New("세미콜론 없음"))
	}
	return letState
}

func ParseReturnStatement() *ast.ReTurnStatement {
	var reState = &ast.ReTurnStatement{}
	return reState
}
