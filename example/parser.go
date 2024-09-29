package main

import (
	"fmt"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/parser"
	"github.com/nam9nine/interpreter/token"
)

func parserLetStatementV1() {
	input := `
let a = 5;
`
	l := lexer.New(input)
	p := parser.New(l)
	tok := token.Token{
		Literal: "let",
		Type:    token.LET,
	}
	st := p.ParseLetStatement(tok)

	fmt.Printf("%v\n", *st.Name)
}

func parseProgramV1() {
	input := `
let a = 5; 
let b = fun(1, 3);
`
	l := lexer.New(input)
	p := parser.New(l)
	sts := p.ParseProgram()
	for _, v := range sts {

		if v == nil {
			fmt.Println("존재하지 않는 예약어")
			return
		}

		if v.TokenLiteral() == "let" {
			fmt.Printf("let 예약어 발견 : %+v\n", v.TokenLiteral())
		} else {
			fmt.Println("다른 예약어 발견")
		}

	}
}

func main() {
	parseProgramV1()
}
