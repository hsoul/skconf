package ast

type Program struct {
	Imports    []ImportStatement
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out string
	for _, imp := range p.Imports {
		out += imp.String() + "\n"
	}
	for _, s := range p.Statements {
		out += s.String() + "\n"
	}
	return out
}
