package ast

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
