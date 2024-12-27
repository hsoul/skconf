package lua

import (
	"github.com/hsoul/skconf/internal/ast"
)

func (l *luaGenerator) generateSkillDef(skill *ast.SkillDef) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("local ")

	l.generateExpression(skill.Name)
	l.buf.WriteString(" = {\n")

	l.indent++
	for i, prop := range skill.Properties {
		l.buf.WriteString(l.indent_str())
		if prop.Key != nil {
			l.generateExpression(prop.Key)
			l.buf.WriteString(" = ")
			if fn, ok := prop.Value.(*ast.FunctionDef); ok {
				fn.Name = prop.Key
			}
		}
		l.generateExpression(prop.Value)
		if i < len(skill.Properties)-1 {
			l.buf.WriteString(",")
		}
		l.buf.WriteString("\n")
	}
	l.indent--

	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("}\n")

	tid := ast.FindPropertyByName("tid", skill.Properties)
	if tid != "" {
		l.skillMap[tid] = skill.Name.Value
	}
}

func (l *luaGenerator) generateStateDef(state *ast.StateDef) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("local ")
	l.generateExpression(state.Name)
	l.buf.WriteString(" = {\n")

	l.indent++
	for i, prop := range state.Properties {
		l.buf.WriteString(l.indent_str())
		if prop.Key != nil {
			l.generateExpression(prop.Key)
			l.buf.WriteString(" = ")
			if fn, ok := prop.Value.(*ast.FunctionDef); ok {
				fn.Name = prop.Key
			}
		}
		l.generateExpression(prop.Value)
		if i < len(state.Properties)-1 {
			l.buf.WriteString(",")
		}
		l.buf.WriteString("\n")
	}
	l.indent--

	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("}\n")

	tid := ast.FindPropertyByName("tid", state.Properties)
	if tid != "" {
		l.stateMap[tid] = state.Name.Value
	}
}
