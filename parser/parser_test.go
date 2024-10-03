package parser

import (
	"fmt"
	"github.com/nam9nine/interpreter/ast"
	"github.com/nam9nine/interpreter/lexer"
	"github.com/nam9nine/interpreter/token"
	"testing"
)

/**
테스트 하고 싶은 것

input -> lexer => 토큰 생성
lexer -(token)-> parser
*/

func TestParseLetStatement(t *testing.T) {
	var correctCount int = 1

	tests := []struct {
		name  string
		input string
		// 기대하는 식별자
		expectedIdentifier string
		// 기대하는 명령문(예약어)
		expectedStatement string
		//기대하는 표현식
		// 곧 추가될 예정

		expectedErrMsg string
		// 오류 코드인 지 성공 코드인 지 판별
		isWrong bool
	}{
		{
			name:           "Invalid Let Statement Missing Identifier",
			input:          `let 5; `,
			expectedErrMsg: fmt.Sprintf("parser error: expected type: %v, Got: %v", token.IDENT, token.INT),
			isWrong:        true,
		},
		{
			name:               "Valid Let statement",
			input:              `let b = 10;`,
			expectedIdentifier: "b",
			isWrong:            false,
		},
		{
			name:           "Invalid Let Statement Missing Assing",
			input:          `let b 5;`,
			expectedErrMsg: fmt.Sprintf("parser error: expected type: %v, Got: %v", token.ASSIGN, token.INT),
			isWrong:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := New(l)
			stmts := p.ParseProgram()
			// 옳은 input 테스트
			if !tt.isWrong {
				if len(stmts) != correctCount {
					t.Fatalf("statements has %d", correctCount)
				}

				for _, st := range stmts {
					//명령문(예약어) 검증
					if st.TokenLiteral() != "let" {
						t.Fatalf("expected statement is let, Got: %s", st.TokenLiteral())
					}

					v, ok := st.(*ast.LetStatement)
					if !ok {
						t.Fatal("not let statement")
					}
					// 식별자 검증
					if v.Name.Value != tt.expectedIdentifier {
						t.Fatalf("Expected identifier: %v, Got: %v", tt.expectedIdentifier, st.TokenLiteral())
					}
					// 식별자 검증
					if v.Name.TokenLiteral() != tt.expectedIdentifier {
						t.Fatalf("Expected identifier: %v, Got: %v", tt.expectedIdentifier, st.TokenLiteral())
					}
					// 표현식 검사 나중에
				}
				// 오류가 있는 input 테스트
			} else {
				// 에러 메세지와 에러 함수 검증
				if p.Errors()[0] == tt.expectedErrMsg {
					t.Log("test complete")
				} else {
					t.Errorf("Expected error message: %q, Got: %q", tt.expectedErrMsg, p.Errors()[0])
				}
			}
		})
	}
}
