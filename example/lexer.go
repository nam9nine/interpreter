package main

import (
	"fmt"
	"github.com/nam9nine/interpreter/token"
)

func LexerExample(t token.Token, input string, i int) (int, token.Token) {
	switch l := input[i]; l {
	case '}':
		i++
	case '{':
		i++
	case ';':
		i++
		t.Type = token.SEMICOLON
		t.Literal = ";"
	case '(':
		t.Type = token.LPAREN
		t.Literal = "("
		i++
	case ')':
		t.Type = token.RPAREN
		t.Literal = ")"
		i++
	case '+':
		t.Type = token.PLUS
		t.Literal = "+"
		i++
	case '=':
		t.Type = token.ASSIGN
		t.Literal = "="
		i++
	case 0x20:
		i++
	default:
		i, t = wordParse(t, input, i)

	}
	return i, t
}

func wordParse(t token.Token, input string, ir int) (int, token.Token) {
	var word []byte
	var tokenT token.TokenType

	for ; ir < len(input); ir++ {
		if isInt(input[ir]) {
			word = append(word, input[ir])
			tokenT = token.INT
		} else {
			break
		}
	}

	for ; ir < len(input); ir++ {
		if isWord(input[ir]) {
			word = append(word, input[ir])
			tokenT = token.IDENT
		} else {
			break
		}
	}

	switch w := string(word); w {
	case "let":
		t.Type = token.LET
		t.Literal = "LET"
	case "fn":
		t.Type = token.FUNCTION
		t.Literal = "fn"
	default:
		t.Type = tokenT
		t.Literal = string(w)
	}
	return ir, t
}

func isWord(l byte) bool {
	return l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z'
}

func isInt(l byte) bool {
	return l >= '0' && l <= '9'
}

func main() {
	input := "let a = 5858;"
	var t token.Token
	var i int = 0
	for i < len(input) {
		i, t = LexerExample(t, input, i)
		fmt.Println("type : ", t.Type)
		fmt.Println("literal : ", t.Literal)
		fmt.Println()
	}
}
