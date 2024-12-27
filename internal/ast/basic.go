package ast

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
