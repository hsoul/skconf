package ast

var _ Node = (*ImportStatement)(nil)
var _ Statement = (*ImportStatement)(nil)

type ImportStatement struct {
	BaseNode
	Value Expression
}

type ElseStatement struct {
	BaseNode
	Condition   Expression // nil for else, non-nil for else-if
	Consequence *CodeBlock
}

type IfStatement struct {
	BaseNode
	Condition    Expression
	Consequence  *CodeBlock
	Alternatives []*ElseStatement
}

type ReturnStatement struct {
	BaseNode
	ReturnValue Expression
}

type CommentStatement struct {
	BaseNode
	Value string
}

type BreakStatement struct { // BreakStatement represents a break statement in a loop
	BaseNode
}

type ContinueStatement struct { // ContinueStatement represents a continue statement in a loop
	BaseNode
}
