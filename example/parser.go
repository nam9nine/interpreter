package main

import (
	"fmt"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/parser"
	"github.com/nam9nine/interpreter/token"
	"os"
)

func parserLetStatementV1() {
	input := `
let a = 5;
`
	l := lexer.New(input)
	tok := token.Token{
		Literal: "let",
		Type:    token.LET,
	}
	st, err := parser.ParseLetStatement(tok, l)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	fmt.Printf("%v\n", *st.Name)
}

func parseProgramV1() {
	input := `
let a = 5;
`
	l := lexer.New(input)
	sts := parser.ParseProgram(l)
	for _, v := range sts {
		fmt.Printf("%+v\n", v.TokenLiteral())
	}
}

func main() {
	parseProgramV1()
}
