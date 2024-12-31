package ast

type TableDef struct {
	BaseNode
	Properties []*PropertyDef
}

type Identifier struct {
	BaseNode
	Value string
}

type Float struct {
	BaseNode
	Value float64
}

type Integer struct {
	BaseNode
	Value int64
}

type String struct {
	BaseNode
	Value string
}

type Boolean struct {
	BaseNode
	Value bool
}

type PrefixExpression struct {
	BaseNode
	Operator string
	Right    Expression
}

type InfixExpression struct {
	BaseNode
	Left     Expression
	Operator string
	Right    Expression
}

type DotExpression struct {
	BaseNode
	Left  Expression
	Right Expression
}

type FunctionCall struct {
	BaseNode
	Function  Expression
	Arguments []Expression
}

type FunctionDef struct {
	BaseNode
	Name       Expression
	Parameters []*Identifier
	Body       *CodeBlock
}
