package ast

import "github.com/lutefd/monkey-lang/token"

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
	Name *Identifier // This is an Identifier to keep things simple and to allow us to later reuse it to represent an identifier as a part of or as a complete expression 
	Value Expression
}


func (ls *LetStatement) statementNode(){}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode(){}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

