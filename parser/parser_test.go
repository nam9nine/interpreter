package parser

import (
	"github.com/nam9nine/interpreter/ast"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
	"testing"
)

func TestParseLetStatement(t *testing.T) {
	//	input1 := `
	//let x = 5;
	//let y = 10;
	//let foobar = 838383;
	//`
	input := `
let a = 5;
`
	// letStatement 구조체
	test := struct {
		expectedName  *ast.Identifier
		expectedToken token.Token
	}{
		expectedName: &ast.Identifier{
			Token: token.Token{
				Literal: "a",
				Type:    token.IDENT,
			},
			Value: "a",
		},
		expectedToken: token.Token{
			Literal: "let",
			Type:    token.LET,
		},
	}
	tok := token.Token{
		Type:    token.LET,
		Literal: "let",
	}

	l := lexer.New(input)
	l.NextToken()
	letState, err := ParseLetStatement(tok, l)
	if err != nil {
		t.Fatal(err)
	} else {
		if letState.Name.Value != test.expectedName.Value {
			t.Fatalf("expected : %+v got : %+v", test.expectedName.Value, letState.Name.Value)
		}
	}
}
