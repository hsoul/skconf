package ast

// VarStatement represents a variable declaration
type VarStatement struct {
	BaseNode
	Name  *Identifier
	Value Expression
}
