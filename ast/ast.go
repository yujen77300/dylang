package ast

import "github.com/yujen77300/dylang/token"

// Node defines an interface for all nodes in the AST.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method as "marker"
}

type Expression interface {
	Node
	expressionNode()
}

// Program node is going to be the root node of every AST our parser produces
type Program struct {
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

// let <identifier> = <expression>;
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()  {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
