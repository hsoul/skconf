package lua

import (
	"github.com/hsoul/skconf/internal/ast"
)

func isSkillProcessFunc(fn *ast.FunctionDef) bool {
	specialFuncs := map[string]bool{
		"XX1": true,
		"XX2": true,
		"XX3": true,
		"XX4": true,
		"XX5": true,
		"YY1": true,
		"YY2": true,
	}
	if identifyer, ok := fn.Name.(*ast.Identifier); ok {
		return specialFuncs[identifyer.Value]
	}
	return false
}

func isSkillFuncCall(call *ast.FunctionCall) bool {
	specialDotLeft := map[string]bool{
		"UE": true,
		"UF": true,
	}
	if dot, ok := call.Function.(*ast.DotExpression); ok {
		if identifyer, ok := dot.Left.(*ast.Identifier); ok {
			return specialDotLeft[identifyer.Value]
		}
	}
	return false
}

func (l *luaGenerator) generateFunctionDef(fn *ast.FunctionDef) {
	l.buf.WriteString("function(")
	if isSkillProcessFunc(fn) {
		l.buf.WriteString("ctx")
		if len(fn.Parameters) > 0 {
			l.buf.WriteString(", ")
		}
	}

	for i, param := range fn.Parameters {
		l.buf.WriteString(param.Value)
		if i < len(fn.Parameters)-1 {
			l.buf.WriteString(", ")
		}
	}
	l.buf.WriteString(")\n")

	l.indent++
	if fn.Body != nil {
		for _, stmt := range fn.Body.Statements {
			l.generateNode(stmt)
		}
	}

	l.indent--
	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("end")
}

func (l *luaGenerator) generateFunctionCall(call *ast.FunctionCall) {
	l.generateExpression(call.Function)
	l.buf.WriteString("(")

	if isSkillFuncCall(call) {
		l.buf.WriteString("ctx")
		if len(call.Arguments) > 0 {
			l.buf.WriteString(", ")
		}
	}

	for i, arg := range call.Arguments {
		l.generateExpression(arg)
		if i < len(call.Arguments)-1 {
			l.buf.WriteString(", ")
		}
	}

	l.buf.WriteString(")")
}
