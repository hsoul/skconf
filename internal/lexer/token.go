package lexer

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF

	// 标识符和字面量
	IDENTIFIER // 标识符
	INTEGER    // 整数
	FLOAT      // 浮点数
	STRING     // 字符串
	BOOLEAN    // true/false
	COMMENT    // 注释

	// 运算符
	ASSIGN    // =
	PLUS      // +
	MINUS     // -
	MULTIPLY  // *
	DIVIDE    // /
	EQ        // ==
	NEQ       // !=
	LT        // <
	GT        // >
	LTE       // <=
	GTE       // >=
	NOT       // not
	AND       // and
	OR        // or
	DOT       // .
	INCREMENT // ++
	DECREMENT // --

	// 分隔符
	COMMA     // ,
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LBRACKET  // [
	RBRACKET  // ]
	SEMICOLON // ;

	// 关键字
	SKILL    // skill
	STATE    // state
	IMPORT   // import
	FUNC     // func
	IF       // if
	ELSE     // else
	RETURN   // return
	TRUE     // true
	FALSE    // false
	VAR      // var
	FOR      // for
	BREAK    // break
	CONTINUE // continue
	RANGE    // range
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case INTEGER:
		return "INTEGER"
	case FLOAT:
		return "FLOAT"
	case STRING:
		return "STRING"
	case BOOLEAN:
		return "BOOLEAN"
	case COMMENT:
		return "COMMENT"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULTIPLY:
		return "MULTIPLY"
	case DIVIDE:
		return "DIVIDE"
	case EQ:
		return "EQ"
	case NEQ:
		return "NEQ"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case LTE:
		return "LTE"
	case GTE:
		return "GTE"
	case NOT:
		return "NOT"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case DOT:
		return "DOT"
	case COMMA:
		return "COMMA"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LBRACKET:
		return "LBRACKET"
	case RBRACKET:
		return "RBRACKET"
	case SEMICOLON:
		return "SEMICOLON"
	case SKILL:
		return "SKILL"
	case STATE:
		return "STATE"
	case IMPORT:
		return "IMPORT"
	case FUNC:
		return "FUNC"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case RETURN:
		return "RETURN"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case VAR:
		return "VAR"
	case FOR:
		return "FOR"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case RANGE:
		return "RANGE"
	case INCREMENT:
		return "INCREMENT"
	case DECREMENT:
		return "DECREMENT"
	default:
		return "UNKNOWN"
	}
}

func (t TokenType) TokenLiteral() string {
	switch t {
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MULTIPLY:
		return "*"
	case DIVIDE:
		return "/"
	case ASSIGN:
		return "="
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case LT:
		return "<"
	case GT:
		return ">"
	case LTE:
		return "<="
	case GTE:
		return ">="
	case NOT:
		return "not"
	case AND:
		return "and"
	case DOT:
		return "."
	case COMMA:
		return ","
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case LBRACKET:
		return "["
	case RBRACKET:
		return "]"
	default:
		return t.String()
	}
}

var keywords = map[string]TokenType{
	"skill":    SKILL,
	"state":    STATE,
	"import":   IMPORT,
	"func":     FUNC,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"true":     TRUE,
	"false":    FALSE,
	"not":      NOT,
	"and":      AND,
	"or":       OR,
	"var":      VAR,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"range":    RANGE,
}

type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}

func (t Token) String() string {
	return fmt.Sprintf("{%-12v, %-16q, %4d, %4d}", t.Type, t.Literal, t.Pos.Line, t.Pos.Column)
}

func centerString(str string, width int) string {
	padding := width - len(str)
	if padding <= 0 {
		return str
	}
	leftPad := padding / 2
	rightPad := padding - leftPad
	return strings.Repeat(" ", leftPad) + str + strings.Repeat(" ", rightPad)
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
