package ast

var _ Node = (*ImportStatement)(nil)
var _ Statement = (*ImportStatement)(nil)

type ExprStmt struct {
	BaseNode
	Expression Expression
}

type ImportStatement struct {
	BaseNode
	Value Expression
}

type VarStatement struct {
	BaseNode
	Name  *Identifier
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

// ForStatement represents a for loop with three forms:
// 1. Classic: for init; condition; post { }
// 2. Condition-only: for condition { }
// 3. Range: for key, value = range array { }
type ForStatement struct {
	BaseNode
	// Classic for and condition-only for
	Init      Statement  // Optional initialization statement
	Condition Expression // Loop condition (required for classic and condition-only)
	Post      Statement  // Optional post statement

	// Range for
	Key        *Identifier // Optional key variable for range
	Value      *Identifier // Value variable for range
	RangeValue Expression  // Expression to range over

	// Common
	Body        *CodeBlock
	IsRangeForm bool // true if this is a range-based for loop
}
