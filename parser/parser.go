package parser

import (
	"BudLang/ast"
	"BudLang/lexer"
	"BudLang/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read the next 2 tokens to set both curToken and nextToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.curToken.Tyoe != token.EOF {
		s := p.parseStatement()
		if s != nil {
			prog.Statements = append(prog.Statements, s)
		}
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseStatement() *ast.Program {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}
