package ast

type PropertyDef struct {
	BaseNode
	Key   Expression
	Value Expression
}

type CodeBlock struct {
	BaseNode
	Statements []Statement
}
