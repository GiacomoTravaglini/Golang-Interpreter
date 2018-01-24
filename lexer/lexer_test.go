// lexer/lexer_test.go
package lexer

import(
	"testing"
	"interpreter/token"
)

func TestNewToken(t* testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y
	};

	let result = add(five, ten);

	!-/*5;
	5 < 10 < 5
	`

	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET , "let"},
		{token.IDENT , "five"},
		{token.ASSIGN , "="},
		{token.INT , "5"},
		{token.SEMICOLON , ";"},
		{token.LET , "let"},
		{token.IDENT , "ten"},
		{token.ASSIGN , "="},
		{token.INT , "10"},
		{token.SEMICOLON , ";"},
		{token.LET , "let"},
		{token.IDENT , "add"},
		{token.ASSIGN , "="},
		{token.FUNCTION , "fn"},
		{token.LPAREN , "("},
		{token.IDENT , "x"},
		{token.COMMA , ","},
		{token.IDENT , "y"},
		{token.RPAREN , ")"},
		{token.LBRACE , "{"},
		{token.IDENT , "x"},
		{token.PLUS , "+"},
		{token.IDENT , "y"},
		{token.RBRACE , "}"},
		{token.SEMICOLON , ";"},
		{token.LET , "let"},
		{token.IDENT , "result"},
		{token.ASSIGN , "="},
		{token.IDENT , "add"},
		{token.LPAREN , "("},
		{token.IDENT , "five"},
		{token.COMMA , ","},
		{token.IDENT , "ten"},
		{token.RPAREN , ")"},
		{token.SEMICOLON , ";"},

		{token.BANG , "!"},
		{token.MINUS , "-"},
		{token.SLASH , "/"},
		{token.ASTERISK , "*"},
		{token.INT , "5"},
		{token.SEMICOLON , ";"},
		{token.INT , "5"},
		{token.LT , "<"},
		{token.INT , "10"},
		{token.LT , "<"},
		{token.INT , "5"},
		{token.EOF , ""},
	}

	l := New(input)

	var tok token.Token
	for i, tt := range tests {
		tok = l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("[%d] Error type: %q, expected: %q",i, tok.Type, tt.expectedType)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("[%d] Error %q, %q",i, tok.Literal, tt.expectedLiteral)
		}
	}
}
