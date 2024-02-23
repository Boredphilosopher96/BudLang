package lexer

import (
	"BudLang/token"
	"go/token"

	"github.com/sqls-server/sqls/token"
)

// TODO support unicode characters in variable names Can do this by changing Lexer.c from byte to rune
// Unicode characters can be multiple bytes long

type Lexer struct {
	input        string
	position     int
	readPosition int
	c            byte
}

func newLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.c = 0
	} else {
		l.c = l.input[l.readPosition]
	}
	l.readPosition = l.position
	l.readPosition += 1
}

func (l *Lexer) NextToken() *token.Token {
	var t token.Token

	switch l.c {
	case "=":
		t = token.NewToken(token.ASSIGN, l.c)

	case "(":
		t = token.NewToken(token.LPAREN, l.c)

	case ")":
		t = token.NewToken(token.RPAREN, l.c)

	case "{":
		t = token.NewToken(token.LBRACE, l.c)

	case "}":
		t = token.NewToken(token.RBRACE, l.c)

	case "+":
		t = token.NewToken(token.ADD, l.c)

	case "/":
		t = token.NewToken(token.DIVIDE, l.c)

	case "*":
		t = token.NewToken(token.MULTIPLY, l.c)

	case ";":
		t = token.NewToken(token.SEMICOLON, l.c)

	case ",":
		t = token.NewToken(token.COMMA, l.c)

	case 0:
		t.Literal = ""
		t.Type = token.EOF
	}

	l.readChar()
	return t
}
