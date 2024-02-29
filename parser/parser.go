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

	for p.curToken.Type != token.EOF {
		s := p.parseStatement()

		// case token.RETURN:
		//     return p.parseReturnStatement()
		// case token.IF:
		//     return p.parseIfStatement()
		// case token.FOR:
		//     return p.parseForStatement()
		// case token.ELSE:
		//     return p.parseElseStatement()
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
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	st := &ast.LetStatement{Token: p.curToken}

	if p.peekToken.Type != token.IDENT {
		return nil
	}
	p.nextToken()

	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	
    if p.peekToken.Type != token.ASSIGN {
		return nil
	}
	p.nextToken()

	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
    return st
}
