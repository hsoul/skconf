package ast

import "github.com/hsoul/skconf/internal/lexer"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

type Expression interface {
	Node
}

type BaseNode struct {
	Token lexer.Token
}

func (b *BaseNode) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BaseNode) String() string {
	return b.Token.Literal
}
