package ast

import "monke/token"

type Node interface {
	TokenLiteral() string
}

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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Statement
}

type ReturnStatement struct {
	Token       token.Token // return
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token       token.Token // first token of the expression
	Expression Expression
}

func (rs *ExpressionStatement) statementNode()       {}
func (rs *ExpressionStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Implement the interface
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
