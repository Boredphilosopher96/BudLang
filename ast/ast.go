package ast

import "BudLang/token"

// Represents a node in the AST
type Node interface {
	TokenLiteral() string
}

// Every line in our language is either a statement or an expression
// So we define statement and expression interfaces and handle them differently

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Value string
	Token token.Token
}

func (ident *Identifier) expressionNode() {}

func (ident *Identifier) TokenLiteral() string { return ident.Token.Literal }

type LetStatement struct {
	Token token.Token
	Value Expression
	Name  *Identifier
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }
