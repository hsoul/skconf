package lua

import (
	"github.com/hsoul/skconf/internal/ast"
)

func (l *luaGenerator) generateTableDef(table *ast.TableDef) {
	l.buf.WriteString("{\n")
	l.indent++

	for i, prop := range table.Properties {
		l.buf.WriteString(l.indent_str())
		if prop.Key != nil {
			switch key := prop.Key.(type) {
			case *ast.Identifier:
				l.buf.WriteString(key.Value)
			case *ast.String:
				l.buf.WriteString("[")
				l.generateExpression(prop.Key)
				l.buf.WriteString("]")
			default:
				l.buf.WriteString("[")
				l.generateExpression(prop.Key)
				l.buf.WriteString("]")
			}
			l.buf.WriteString(" = ")
		}

		switch val := prop.Value.(type) { // Handle nested tables with proper indentation
		case *ast.TableDef:
			l.generateTableDef(val)
		case *ast.FunctionDef:
			l.generateFunctionDef(val)
		default:
			l.generateExpression(prop.Value)
		}

		if i < len(table.Properties)-1 {
			l.buf.WriteString(",")
		}
		l.buf.WriteString("\n")
	}

	l.indent--
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("}")
}
