package ast

import (
	"BudLang/token"
	"bytes"
)

// Represents a node in the AST
type Node interface {
	TokenLiteral() string
    String() string
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

func (p *Program) String() string {
    var output bytes.Buffer
    for s, stmt := range p.Statements {
        output.WriteString(stmt.String())
        if s < len(p.Statements) - 1 {
            output.WriteString("\n")
        }
    }
    return output.String()
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

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

func (l *LetStatement) String() string {
    var output bytes.Buffer
    output.WriteString(l.TokenLiteral()+ " ")
    output.WriteString(l.Name.Value)
    output.WriteString(" = ")
    if l.Value != nil {
        output.WriteString(l.Value.String())
    }
    output.WriteString(";")
    return output.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode()      {}
func (r *ReturnStatement) TokenLiteral() string { return r.Token.Literal }

func (r *ReturnStatement) String() string {
    var output bytes.Buffer
    output.WriteString(r.TokenLiteral() + " ")
    if r.ReturnValue != nil {
        output.WriteString(r.ReturnValue.String())
    }
    output.WriteString(";")
    return output.String()
}

type ExpressionStatement struct {
    Token token.Token
    ExpressionValue Expression
}

func (r *ExpressionStatement) statementNode()      {}
func (r *ExpressionStatement) TokenLiteral() string { return r.Token.Literal }

func (r *ExpressionStatement) String() string {
    if r.ExpressionValue != nil {
        return r.ExpressionValue.String()
    }
    return ""
}

type (
    prefixParserFunc func() Expression
    infixParserFunc func(Expression) Expression
)
