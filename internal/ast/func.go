package ast

type CodeBlock struct {
	BaseNode
	Statements []Statement
}

type ExprStmt struct {
	BaseNode
	Expression Expression
}

type FunctionDef struct {
	BaseNode
	Parameters []*Identifier
	Body       *CodeBlock

	// generator props
	Name Expression
}
