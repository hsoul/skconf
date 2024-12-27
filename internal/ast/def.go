package ast

type Block struct {
	BaseNode
	Statements []Statement
}

type SkillDef struct {
	BaseNode
	Name       *Identifier
	Properties []*PropertyDef
}

type StateDef struct {
	BaseNode
	Name       *Identifier
	Properties []*PropertyDef
}
