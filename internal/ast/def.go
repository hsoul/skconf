package ast

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
