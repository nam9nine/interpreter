package main

import "fmt"

func LexerExample() {
	input := "let ajjjj = 5;"

	for i := 0; i < len(input); i++ {
		l := input[i]
		switch l {
		case '}':
			{
				fmt.Println('}')

			}
		case '{':
			fmt.Println('{')
		case ';':
			fmt.Println(';')
		case '(':
			fmt.Println('(')
		case ')':
			fmt.Println(')')
		case '+':
			fmt.Println('+')
		case '=':
			fmt.Println('=')
		case 0x20:
			continue
		default:

			i = wordParse(input, i)
		}

	}
}

func wordParse(input string, ir int) int {

	var word []byte
	var start = ir
	for isWord(input[start]) {
		word = append(word, input[start])
		start = start + 1
	}
	switch w := string(word); w {
	case "let":
		fmt.Println("let")
	case "fn":
		fmt.Println("fn")
	default:
		fmt.Println(w)
	}
	return start
}

func isWord(l byte) bool {
	return l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z'
}

func main() {
	LexerExample()
}
