package lua

import (
	"github.com/hsoul/skconf/internal/ast"
)

func (l *luaGenerator) generateForStatement(stmt *ast.ForStatement) {
	l.buf.WriteString(l.indent_str())

	isBraced := false

	if stmt.IsRangeForm {
		l.buf.WriteString("for ")
		if stmt.Key != nil {
			l.generateExpression(stmt.Key)
			l.buf.WriteString(", ")
		}
		l.generateExpression(stmt.Value)
		l.buf.WriteString(" in pairs(")
		l.generateExpression(stmt.RangeValue)
		l.buf.WriteString(") do\n")
	} else if isNumericForLoop(stmt) {
		l.buf.WriteString("for ")
		ok := l.generateNumericForParams(stmt)
		if !ok {
			return
		}
		l.buf.WriteString(" do\n")
	} else {
		if stmt.Init != nil {
			if varStmt, ok := stmt.Init.(*ast.VarStatement); ok {
				isBraced = true
				l.indent++
				l.buf.WriteString("do\n")
				l.buf.WriteString(l.indent_str())
				l.buf.WriteString("local ")
				l.generateExpression(varStmt.Name)
				l.buf.WriteString(" = ")
				l.generateExpression(varStmt.Value)
				l.buf.WriteString("\n")
				l.buf.WriteString(l.indent_str())
				l.buf.WriteString("while ")
				l.generateExpression(stmt.Condition)
				l.buf.WriteString(" do\n")
			}
		} else {
			l.buf.WriteString("while ")
			l.generateExpression(stmt.Condition)
			l.buf.WriteString(" do\n")
		}
	}

	l.indent++
	if stmt.Body != nil {
		for _, s := range stmt.Body.Statements {
			l.generateNode(s)
		}
	}

	if !stmt.IsRangeForm && !isNumericForLoop(stmt) && stmt.Post != nil {
		if exprStmt, ok := stmt.Post.(*ast.ExprStmt); ok {
			l.buf.WriteString(l.indent_str())
			l.generateExpression(exprStmt.Expression)
			l.buf.WriteString("\n")
		}
	}
	l.indent--

	l.buf.WriteString(l.indent_str())
	l.buf.WriteString("end\n")
	if isBraced {
		l.indent--
		l.buf.WriteString(l.indent_str())
		l.buf.WriteString("end\n")
	}
}

func isNumericForLoop(stmt *ast.ForStatement) bool {
	initVar, initVal := extractInitStatement(stmt)
	if initVar == nil || initVal == nil {
		return false
	}

	condVar, condOp, condVal := extractCondition(stmt.Condition)
	if condVar == nil || condVal == nil {
		return false
	}

	if condOp != "<=" && condOp != ">=" {
		return false
	}

	if initVar.String() != condVar.String() {
		return false
	}

	postVar, _, postVal := extractPostStatement(stmt.Post)
	if postVar == nil || postVal == nil {
		return false
	}

	if postVar.String() != initVar.String() {
		return false
	}

	return true
}

func (l *luaGenerator) generateNumericForParams(stmt *ast.ForStatement) bool {
	var start, end, step ast.Expression

	initVar, start := extractInitStatement(stmt)
	if initVar == nil {
		return false
	}

	condVar, condOp, condVal := extractCondition(stmt.Condition)
	if condVar == nil {
		return false
	}

	postVar, postOp, postVal := extractPostStatement(stmt.Post)
	if postVar == nil {
		return false
	}

	switch condOp {
	case "<=":
		end = condVal
	case ">=":
		end = condVal
	}

	step = postVal

	l.generateExpression(initVar)
	l.buf.WriteString(" = ")
	l.generateExpression(start)
	l.buf.WriteString(", ")
	l.generateExpression(end)
	if step != nil {
		isWrite := true
		if value, ok := step.(*ast.Integer); ok {
			if postOp == "+" && value.Value == 1 {
				isWrite = false
			}
		}

		if isWrite {
			l.buf.WriteString(", ")
			if postOp == "-" {
				l.buf.WriteString("-")
			}
			l.generateExpression(step)
		}
	}

	return true
}

func extractInitStatement(stmt *ast.ForStatement) (varExpr, valExpr ast.Expression) {
	if stmt.Init == nil {
		return nil, nil
	}

	if varStmt, ok := stmt.Init.(*ast.VarStatement); ok {
		return varStmt.Name, varStmt.Value
	}

	return nil, nil
}

func extractCondition(expr ast.Expression) (varExpr ast.Expression, operator string, valExpr ast.Expression) {
	if infixExpr, ok := expr.(*ast.InfixExpression); ok {
		return infixExpr.Left, infixExpr.Operator, infixExpr.Right
	}
	return nil, "", nil
}

func extractPostStatement(stmt ast.Statement) (varExpr ast.Expression, operator string, valExpr ast.Expression) {
	if stmt == nil {
		return nil, "", nil
	}

	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return nil, "", nil
	}

	assignInfix, ok := exprStmt.Expression.(*ast.InfixExpression)
	if !ok {
		return nil, "", nil
	}

	if assignInfix.Operator != "=" {
		return nil, "", nil
	}

	operInfix, ok := assignInfix.Right.(*ast.InfixExpression)
	if !ok {
		return nil, "", nil
	}

	if !ast.IsSameIdentifier(assignInfix.Left, operInfix.Left) {
		return nil, "", nil
	}

	switch operInfix.Operator {
	case "+", "-":
		return assignInfix.Left, operInfix.Operator, operInfix.Right
	default:
		return nil, "", nil
	}
}
