package ast

type Number interface {
	number()
}

func (f *Float) number()   {}
func (i *Integer) number() {}

type Value interface {
	value()
}

func (f *Float) value()       {}
func (i *Integer) value()     {}
func (t *TableDef) value()    {}
func (f *FunctionDef) value() {}
func (i *Identifier) value()  {}
func (s *String) value()      {}
func (b *Boolean) value()     {}

type Statement interface {
	Node
	statement()
}

func (e *ExprStmt) statement()          {}
func (i *ImportStatement) statement()   {}
func (v *VarStatement) statement()      {}
func (f *FunctionDef) statement()       {}
func (i *IfStatement) statement()       {}
func (r *ReturnStatement) statement()   {}
func (c *CommentStatement) statement()  {}
func (b *BreakStatement) statement()    {}
func (c *ContinueStatement) statement() {}
func (s *SkillDef) statement()          {}
func (s *StateDef) statement()          {}
func (f *ForStatement) statement()      {}

type Expression interface {
	Node
	expression()
}

func (f *Float) expression()            {}
func (i *Integer) expression()          {}
func (t *TableDef) expression()         {}
func (f *FunctionDef) expression()      {}
func (i *Identifier) expression()       {}
func (s *String) expression()           {}
func (b *Boolean) expression()          {}
func (p *PrefixExpression) expression() {}
func (i *InfixExpression) expression()  {}
func (d *DotExpression) expression()    {}
func (b *FunctionCall) expression()     {}
