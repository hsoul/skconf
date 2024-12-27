package ast

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

func (d *DotExpression) expressionNode() {}

type FunctionCall struct {
	BaseNode
	Function  Expression
	Arguments []Expression
}
