package repl

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
	"io"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) error {
	scan := bufio.NewScanner(r)
	for {
		fmt.Fprintf(w, PROMPT)
		if !scan.Scan() {
			if err := scan.Err(); err != nil {
				return errors.New("input err")
			} else {
				fmt.Fprintf(w, "EOF발생 REPL을 종료합니다\n")
				return nil
			}
		}
		line := lexer.New(scan.Text())
		for tok := line.NextToken(); tok.Type != token.EOF; tok = line.NextToken() {
			fmt.Fprintf(w, "Type : %+v\n", tok)
		}
	}
}
