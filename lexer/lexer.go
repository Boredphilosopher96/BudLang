package lexer

import (
	"BudLang/token"
)

// TODO support unicode characters in variable names Can do this by changing Lexer.c from byte to rune
// Unicode characters can be multiple bytes long

type Lexer struct {
	input        string
	position     int
	readPosition int
	c            byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.c = 0
	} else {
		l.c = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.eatWhiteSpace()

	switch l.c {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			t = token.NewToken(token.ASSIGN, l.c)
		}
	case '(':
		t = token.NewToken(token.LPAREN, l.c)

	case ')':
		t = token.NewToken(token.RPAREN, l.c)

	case '{':
		t = token.NewToken(token.LBRACE, l.c)

	case '}':
		t = token.NewToken(token.RBRACE, l.c)

	case '+':
		t = token.NewToken(token.PLUS, l.c)

	case '/':
		t = token.NewToken(token.SLASH, l.c)

	case '-':
		t = token.NewToken(token.MINUS, l.c)

	case '*':
		t = token.NewToken(token.ASTERISK, l.c)

	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			t = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			t = token.NewToken(token.BANG, l.c)
		}
	case '<':
		t = token.NewToken(token.LT, l.c)

	case '>':
		t = token.NewToken(token.GT, l.c)

	case ';':
		t = token.NewToken(token.SEMICOLON, l.c)

	case ',':
		t = token.NewToken(token.COMMA, l.c)

	case 0:
		t.Literal = ""
		t.Type = token.EOF

	default:
		if isAllowedChar(l.c) {
			t.Literal = l.multiCharToken()
			t.Type = token.LookupKeywords(t.Literal)
			return t
		} else if isAllowedDigit(l.c) {
			t.Literal = l.getDigits()
			t.Type = token.INT
			return t
		} else {
			t = token.NewToken(token.ILLEGAL, l.c)
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) multiCharToken() string {
	spos := l.position
	for isAllowedChar(l.c) {
		l.readChar()
	}
	return l.input[spos:l.position]
}

func (l *Lexer) getDigits() string {
	spos := l.position
	for isAllowedDigit(l.c) {
		l.readChar()
	}
	return l.input[spos:l.position]
}

func isAllowedChar(c byte) bool {
	return (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'a') || c == '_'
}

func isAllowedDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) eatWhiteSpace() {
	for isWhiteSpace(l.c) {
		l.readChar()
	}
}

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}
