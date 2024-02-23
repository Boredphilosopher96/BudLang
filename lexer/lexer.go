package lexer

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
