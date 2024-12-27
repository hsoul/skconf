package lua

import (
	"github.com/hsoul/skconf/internal/ast"
)

func (l *luaGenerator) generateVarStatement(stmt *ast.VarStatement) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("local ")
	l.generateExpression(stmt.Name)
	l.buf.WriteString(" = ")
	l.generateExpression(stmt.Value)
	l.buf.WriteString("\n")
}

func (l *luaGenerator) generateImportStatement(stmt *ast.ImportStatement) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("require ")
	l.generateExpression(stmt.Value)
}

func (l *luaGenerator) generateIfStatement(stmt *ast.IfStatement) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("if ")
	l.generateExpression(stmt.Condition)
	l.buf.WriteString(" then\n")

	l.indent++
	if stmt.Consequence != nil {
		for _, s := range stmt.Consequence.Statements {
			l.generateNode(s)
		}
	}
	l.indent--

	for _, alt := range stmt.Alternatives { // Handle else-if and else chains
		l.buf.WriteString(l.indent_str())
		if alt.Condition != nil {
			l.buf.WriteString("elseif ")
			l.generateExpression(alt.Condition)
			l.buf.WriteString(" then\n")
		} else {
			l.buf.WriteString("else\n")
		}

		l.indent++
		if alt.Consequence != nil {
			for _, s := range alt.Consequence.Statements {
				l.generateNode(s)
			}
		}
		l.indent--
	}

	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("end\n")
}

func (l *luaGenerator) generateReturnStatement(stmt *ast.ReturnStatement) {
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("return ")
	if stmt.ReturnValue != nil {
		l.generateExpression(stmt.ReturnValue)
	}
	l.buf.WriteString("\n")
}
