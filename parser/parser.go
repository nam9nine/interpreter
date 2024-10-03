package parser

import (
	"fmt"
	"github.com/nam9nine/interpreter/ast"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
)

type Parser struct {
	lex         *lexer.Lexer
	curToken    token.Token
	errMessages []string
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		lex:         l,
		curToken:    l.NextToken(),
		errMessages: []string{},
	}
}

func (p *Parser) NextToken() token.Token {
	tok := p.lex.NextToken()
	p.curToken = tok
	return p.curToken
}

// 반환값 program 구조체 추가
func (p *Parser) ParseProgram() []ast.Statement {
	// 나중에 Program 구조체로 재정의해야됨
	var statements []ast.Statement

	for p.curToken.Type != token.EOF {
		st := p.ParseStatement()
		if st != nil {
			statements = append(statements, st)
		}
		p.NextToken()
	}
	return statements
}

func (p *Parser) ParseStatement() ast.Statement {
	var statement ast.Statement

	tp := p.curToken.Type
	switch tp {
	case token.LET:
		statement = p.ParseLetStatement()
		if t, ok := statement.(*ast.LetStatement); ok {
			if t == nil {
				return nil
			}
		}
	case token.RETURN:
		statement = p.ParseReturnStatement()
		if t, ok := statement.(*ast.ReturnStatement); ok {
			if t == nil {
				return nil
			}
		}
	default:
		return nil
	}
	return statement
}

func (p *Parser) ParseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	} else {
		// let AST에 식별자 추가
		stmt.Name = &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	//표현식 건너뛰기 - 미완성
	for p.curToken.Type != token.SEMICOLON {
		p.NextToken()
	}
	return stmt
}

func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	var stmt = &ast.ReturnStatement{}
	return stmt
}

func (p *Parser) expectPeek(tType token.TokenType) bool {
	tok := p.NextToken()
	if tok.Type != tType {
		p.errMessages = append(p.errMessages,
			fmt.Sprintf("parser error: expected type: %v, Got: %v", tType, tok.Type))
		return false
	}
	return true
}

func (p *Parser) Errors() []string {
	return p.errMessages
}
