package ast

import "github.com/nam9nine/interpreter/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	// 식별자, 대입 연산자, 표혆식
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

type ReTurnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReTurnStatement) statementNode() {

}

func (r *ReTurnStatement) TokenLiteral() string {
	return r.Token.Literal
}
