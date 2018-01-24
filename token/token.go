// token/token.go
package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENT = "IDENT"
	INT = "INT"
	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	FUNCTION = "fn"
	LET = "let"
)

var keywords = map[string]TokenType{
	"let" : LET,
	"fn" : FUNCTION,
}

func LookupIdent(ident string) TokenType {
	if ttype, ok := keywords[ident]; ok {
		return ttype
	}
	return IDENT
}
