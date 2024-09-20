package main

import (
	"github.com/nam9nine/interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
	}
	for i, v := range tests {
		var t token.Token

	}
}
