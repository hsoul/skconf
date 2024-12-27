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
