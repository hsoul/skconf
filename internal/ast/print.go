package ast

import (
	"fmt"
	"strings"
)

func PrintTree(node Node) string {
	return printTreeWithIndent(node, "", true)
}

func printTreeWithIndent(node Node, prefix string, isLast bool) string {
	if node == nil {
		return ""
	}

	var sb strings.Builder

	marker := "└── "
	if !isLast {
		marker = "├── "
	}

	sb.WriteString(prefix)
	sb.WriteString(marker)
	sb.WriteString(nodeToString(node))
	sb.WriteString("\n")

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	children := getChildren(node)
	for i, child := range children {
		isLastChild := i == len(children)-1
		sb.WriteString(printTreeWithIndent(child, childPrefix, isLastChild))
	}

	return sb.String()
}

func nodeToString(node Node) string {
	if labeled, ok := node.(*labeledNode); ok {
		return labeled.String()
	}

	switch n := node.(type) {
	case *Program:
		return fmt.Sprintf("Program { Imports: %d, Statements: %d }", len(n.Imports), len(n.Statements))
	case *SkillDef:
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("SkillDef"))
		return sb.String()
	case *StateDef:
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("StateDef"))
		return sb.String()
	case *PropertyDef:
		return fmt.Sprintf("PropertyDef { Key: %s }", n.Key)
	case *FunctionDef:
		params := make([]string, len(n.Parameters))
		for i, p := range n.Parameters {
			params[i] = p.Value
		}
		return fmt.Sprintf("FunctionDef")
	case *Identifier:
		return fmt.Sprintf("Identifier { Value: %s }", n.Value)
	case *Integer:
		return fmt.Sprintf("Integer { Value: %d }", n.Value)
	case *Float:
		return fmt.Sprintf("Float { Value: %f }", n.Value)
	case *String:
		return fmt.Sprintf("String { Value: %s }", n.Value)
	case *Boolean:
		return fmt.Sprintf("Boolean { Value: %v }", n.Value)
	case *CodeBlock:
		return fmt.Sprintf("CodeBlock")
	case *IfStatement:
		return "IfStatement"
	case *ReturnStatement:
		return "ReturnStatement"
	case *ImportStatement:
		return fmt.Sprintf("ImportStatement { Path: %s }", n.Value)
	case *CommentStatement:
		return fmt.Sprintf("CommentStatement { Text: %s }", n.Value)
	case *InfixExpression:
		return fmt.Sprintf("InfixExpression")
	case *PrefixExpression:
		return fmt.Sprintf("PrefixExpression")
	case *ExprStmt:
		return "ExpressionStatement"
	case *DotExpression:
		return "DotExpression"
	case *FunctionCall:
		return fmt.Sprintf("FunctionCall { Args: %d }", len(n.Arguments))
	case *labeledNode:
		return n.String()
	case *ForStatement:
		if n.IsRangeForm {
			return fmt.Sprintf("ForStatement (Range)")
		}
		return fmt.Sprintf("ForStatement (Classic)")
	default:
		return fmt.Sprintf("%T", node)
	}
}

func getChildren(node Node) []Node {
	if labeled, ok := node.(*labeledNode); ok {
		return getChildren(labeled.node)
	}

	var children []Node

	switch n := node.(type) {
	case *Program:
		for i := range n.Imports {
			children = append(children, &labeledNode{"import", &n.Imports[i]})
		}
		for _, stmt := range n.Statements {
			children = append(children, &labeledNode{"statement", stmt})
		}
	case *SkillDef:
		children = append(children, &labeledNode{"name", n.Name})
		for _, prop := range n.Properties {
			if prop.Key != nil {
				children = append(children, &labeledNode{fmt.Sprintf("property '%s'", prop.Key), prop.Value})
			} else {
				children = append(children, &labeledNode{fmt.Sprintf("property "), prop.Value})
			}
		}
	case *StateDef:
		children = append(children, &labeledNode{"name", n.Name})
		for _, prop := range n.Properties {
			children = append(children, &labeledNode{fmt.Sprintf("property '%s'", prop.Key), prop.Value})
		}
	case *PropertyDef:
		children = append(children, &labeledNode{"key", n.Key})
		children = append(children, &labeledNode{fmt.Sprintf("value '%s'", n.Value), n.Value})
	case *FunctionDef:
		for _, param := range n.Parameters {
			children = append(children, &labeledNode{"parameter", param})
		}
		if n.Body != nil {
			children = append(children, &labeledNode{"body", n.Body})
		}
	case *CodeBlock:
		for _, stmt := range n.Statements {
			children = append(children, &labeledNode{"statement", stmt})
		}
	case *IfStatement:
		if n.Consequence != nil {
			ifCondition := &labeledNode{"if", &labeledNode{"condition", n.Condition}}
			compositionNode := &compositionNode{node: make([]Node, 0)}
			compositionNode.node = append(compositionNode.node, ifCondition, n.Consequence)
			children = append(children, compositionNode)
		} else {
			compositionNode := &compositionNode{node: make([]Node, 0)}
			compositionNode.node = append(compositionNode.node, &labeledNode{"if", &labeledNode{"condition", n.Condition}})
			children = append(children, compositionNode)
		}
		for _, alt := range n.Alternatives {
			children = append(children, alt)
		}
	case *ElseStatement:
		if n.Condition != nil && n.Consequence != nil {
			children = append(children, &labeledNode{"else if", &labeledNode{"condition", n.Condition}}, n.Consequence)
		} else if (n.Condition == nil) && (n.Consequence != nil) {
			children = append(children, &labeledNode{"else", nil}, n.Consequence)
		}
	case *ReturnStatement:
		if n.ReturnValue != nil {
			children = append(children, &labeledNode{"value", n.ReturnValue})
		}
	case *FunctionCall:
		children = append(children, &labeledNode{"function", n.Function})
		for i, arg := range n.Arguments {
			children = append(children, &labeledNode{fmt.Sprintf("arg[%d]", i), arg})
		}
	case *DotExpression:
		children = append(children, &labeledNode{"left", n.Left})
		children = append(children, &labeledNode{"right", n.Right})
	case *ExprStmt:
		children = append(children, &labeledNode{"expression", n.Expression})
	case *TableDef:
		for _, prop := range n.Properties {
			if prop.Key != nil {
				children = append(children, &labeledNode{fmt.Sprintf("property '%s'", prop.Key), prop.Value})
			} else {
				children = append(children, &labeledNode{fmt.Sprintf("property "), prop.Value})
			}
		}
	case *InfixExpression:
		children = append(children, &labeledNode{"left", n.Left})
		children = append(children, &labeledNode{fmt.Sprintf("operator '%s'", n.Operator), nil})
		children = append(children, &labeledNode{"right", n.Right})
	case *ForStatement:
		if n.IsRangeForm {
			if n.Key != nil {
				children = append(children, &labeledNode{"key", n.Key})
			}
			children = append(children, &labeledNode{"value", n.Value})
			children = append(children, &labeledNode{"range", n.RangeValue})
		} else {
			if n.Init != nil {
				children = append(children, &labeledNode{"init", n.Init})
			}
			if n.Condition != nil {
				children = append(children, &labeledNode{"condition", n.Condition})
			}
			if n.Post != nil {
				children = append(children, &labeledNode{"post", n.Post})
			}
		}
		if n.Body != nil {
			children = append(children, &labeledNode{"body", n.Body})
		}
	case *PrefixExpression:
		children = append(children, &labeledNode{fmt.Sprintf("operator '%s'", n.Operator), nil})
		children = append(children, &labeledNode{"right", n.Right})
	case *ImportStatement:
		children = append(children, &labeledNode{fmt.Sprintf("exp"), n.Value})
	case *VarStatement:
		children = append(children, &labeledNode{fmt.Sprintf("var"), n.Name})
		children = append(children, &labeledNode{fmt.Sprintf("exp"), n.Value})
	case *labeledNode:
		children = append(children, n.node)
	case *compositionNode:
		for _, child := range n.node {
			children = append(children, child)
		}
	}
	return children
}

type labeledNode struct {
	label string
	node  Node
}

func (n *labeledNode) TokenLiteral() string { return n.node.TokenLiteral() }
func (n *labeledNode) String() string {
	if n.node == nil {
		return fmt.Sprintf("%s", n.label)
	}
	return fmt.Sprintf("%s: %s", n.label, nodeToString(n.node))
}

type compositionNode struct {
	BaseNode
	node []Node
}
