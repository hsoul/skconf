package ast

type PropertyDef struct {
	BaseNode
	Key   Expression
	Value Expression
}
