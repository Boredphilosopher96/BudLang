package parser

import (
	"BudLang/ast"
	"BudLang/lexer"
	"BudLang/token"
    "fmt"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	// Read the next 2 tokens to set both curToken and nextToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		s := p.parseStatement()
		if s != nil {
			prog.Statements = append(prog.Statements, s)
		}
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	st := &ast.LetStatement{Token: p.curToken}

	if !p.expectToken(token.IDENT) {
		return nil
	}

	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectToken(token.ASSIGN) {
		return nil
	}

	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return st
}

func (p *Parser) expectToken(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	st := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for p.peekToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return st
}
