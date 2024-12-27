package lua

import (
	"fmt"
	"strings"

	"github.com/hsoul/skconf/internal/ast"
)

func (l *luaGenerator) generateExpression(exp ast.Expression) {
	if exp == nil {
		return
	}

	switch n := exp.(type) {
	case *ast.Identifier:
		l.buf.WriteString(strings.ReplaceAll(n.Value, "-", "_"))
	case *ast.Integer:
		l.buf.WriteString(fmt.Sprintf("%d", n.Value))
	case *ast.Float:
		l.buf.WriteString(fmt.Sprintf("%.1f", n.Value))
	case *ast.String:
		l.buf.WriteString(fmt.Sprintf("%q", n.Value))
	case *ast.Boolean:
		if n.Value {
			l.buf.WriteString("true")
		} else {
			l.buf.WriteString("false")
		}
	case *ast.InfixExpression:
		l.generateInfixExpression(n)
	case *ast.PrefixExpression:
		l.generatePrefixExpression(n)
	case *ast.FunctionCall:
		l.generateFunctionCall(n)
	case *ast.DotExpression:
		l.generateDotExpression(n)
	case *ast.TableDef:
		l.generateTableDef(n)
	case *ast.FunctionDef:
		l.generateFunctionDef(n)
	default:
		l.buf.WriteString(fmt.Sprintf("-- Unhandled expression type: %T", exp))
	}
}

func (l *luaGenerator) generateInfixExpression(exp *ast.InfixExpression) {
	// 生成左表达式
	l.generateSubExpression(exp.Left, exp, true)

	l.buf.WriteString(" ")

	// 转换操作符为 Lua 语法
	switch exp.Operator {
	case "and":
		l.buf.WriteString("and")
	case "or":
		l.buf.WriteString("or")
	case "==":
		l.buf.WriteString("==")
	case "!=":
		l.buf.WriteString("~=")
	case "<":
		l.buf.WriteString("<")
	case ">":
		l.buf.WriteString(">")
	case "<=":
		l.buf.WriteString("<=")
	case ">=":
		l.buf.WriteString(">=")
	case "+":
		l.buf.WriteString("+")
	case "-":
		l.buf.WriteString("-")
	case "*":
		l.buf.WriteString("*")
	case "/":
		l.buf.WriteString("/")
	default:
		l.buf.WriteString(exp.Operator)
	}

	l.buf.WriteString(" ")

	// 生成右表达式
	l.generateSubExpression(exp.Right, exp, false)
}

func (l *luaGenerator) generateSubExpression(sub ast.Expression, parent *ast.InfixExpression, isLeft bool) {
	if subInfix, ok := sub.(*ast.InfixExpression); ok {
		needParens := needParentheses(subInfix, parent, isLeft)
		if needParens {
			l.buf.WriteString("(")
		}
		l.generateInfixExpression(subInfix)
		if needParens {
			l.buf.WriteString(")")
		}
	} else {
		l.generateExpression(sub)
	}
}

func needParentheses(sub *ast.InfixExpression, parent *ast.InfixExpression, isLeft bool) bool {
	subPrec := operatorPrecedence(sub.Operator)
	parentPrec := operatorPrecedence(parent.Operator)

	if subPrec < parentPrec {
		return true
	}

	if subPrec == parentPrec { // 对于同优先级，需考虑结合性
		switch parent.Operator {
		case "and", "or":
			if !isLeft {
				return true
			}
		}
	}

	return false
}

func (l *luaGenerator) generateDotExpression(exp *ast.DotExpression) {
	l.generateExpression(exp.Left)
	l.buf.WriteString(".")
	if id, ok := exp.Right.(*ast.Identifier); ok {
		l.buf.WriteString(id.Value)
	} else {
		l.generateExpression(exp.Right)
	}
}

func (l *luaGenerator) generatePrefixExpression(exp *ast.PrefixExpression) {
	switch exp.Operator {
	case "not":
		l.buf.WriteString("not ")
		l.generateExpression(exp.Right)
	default:
		l.buf.WriteString(exp.Operator)
		l.buf.WriteString("(")
		l.generateExpression(exp.Right)
		l.buf.WriteString(")")
	}
}

func (l *luaGenerator) generateExpressionStatement(stmt *ast.ExprStmt) {
	l.buf.WriteString(l.indent_str())
	l.generateExpression(stmt.Expression)
}

func operatorPrecedence(op string) int {
	switch op {
	case "^":
		return 7
	case "*", "/", "%":
		return 6
	case "+", "-":
		return 5
	case "==", "~=", "<", ">", "<=", ">=":
		return 4
	case "and":
		return 3
	case "or":
		return 2
	default:
		return 1
	}
}
