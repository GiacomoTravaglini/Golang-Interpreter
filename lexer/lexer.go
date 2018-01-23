// lexer/lexer.go
package lexer

import(
	"interpreter/token"
)

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input : input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
		case '=':
			tok = NewToken(token.ASSIGN, l.ch)
		case '+':
			tok = NewToken(token.PLUS, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}