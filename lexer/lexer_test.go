// lexer/lexer_test.go
package lexer

import(
	"testing"
	"interpreter/token"
)

func TestNewToken(t* testing.T) {
	input := `=+`
	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN , "="},
		{token.PLUS , "+"},
	}

	l := New(input)

	var tok token.Token
	for _, tt := range tests {
		tok = l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("Error")
		}
	}
}
