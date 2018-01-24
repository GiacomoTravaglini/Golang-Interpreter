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
	MINUS = "-"
	SLASH = "/"
	ASTERISK = "*"
	BANG = "!"

	LT = "<"
	GT = ">"

	COMMA = ","
	SEMICOLON = ";"

	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"

	FUNCTION = "fn"
	LET = "let"
	IF = "if"
	ELSE = "else"
	RETURN = "return"
	TRUE = "true"
	FALSE = "false"
)

var keywords = map[string]TokenType{
	"let" : LET,
	"fn" : FUNCTION,
	"if" : IF,
	"else" : ELSE,
	"return" : RETURN,
	"true" : TRUE,
	"false" : FALSE,
}

func LookupIdent(ident string) TokenType {
	if ttype, ok := keywords[ident]; ok {
		return ttype
	}
	return IDENT
}
