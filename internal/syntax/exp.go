package syntax

import (
	"fmt"

	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

const (
	_           int = iota
	LOWEST          // 最低优先级
	ASSIGN          // =
	OR              // or
	AND             // and
	EQUALS          // == or !=
	LESSGREATER     // > or < or >= or <=
	SUM             // +
	PRODUCT         // *
	PREFIX          // -X or !X
	CALL            // myFunction(X)
	DOT             // module.function
)

var precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NEQ:      EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.LTE:      LESSGREATER,
	lexer.GTE:      LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.MULTIPLY: PRODUCT,
	lexer.DIVIDE:   PRODUCT,
	lexer.LPAREN:   CALL,
	lexer.DOT:      DOT,
	lexer.AND:      AND,
	lexer.OR:       OR,
	lexer.ASSIGN:   ASSIGN,
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPrecedenceFuncs() {
	p.registerPrefix(lexer.IDENTIFIER, p.parseIdentifier) // Register prefix parse functions
	p.registerPrefix(lexer.INTEGER, p.parseInteger)
	p.registerPrefix(lexer.FLOAT, p.parseFloat)
	p.registerPrefix(lexer.STRING, p.parseString)
	p.registerPrefix(lexer.TRUE, p.parseBoolean)
	p.registerPrefix(lexer.FALSE, p.parseBoolean)
	p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(lexer.LBRACE, p.parseTable)
	p.registerPrefix(lexer.FUNC, p.parseFunction)
	p.registerPrefix(lexer.NOT, p.parsePrefixExpression)
	p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)

	p.registerInfix(lexer.PLUS, p.parseInfixExpression) // Register infix parse functions
	p.registerInfix(lexer.MINUS, p.parseInfixExpression)
	p.registerInfix(lexer.MULTIPLY, p.parseInfixExpression)
	p.registerInfix(lexer.DIVIDE, p.parseInfixExpression)
	p.registerInfix(lexer.EQ, p.parseInfixExpression)
	p.registerInfix(lexer.NEQ, p.parseInfixExpression)
	p.registerInfix(lexer.LT, p.parseInfixExpression)
	p.registerInfix(lexer.GT, p.parseInfixExpression)
	p.registerInfix(lexer.LTE, p.parseInfixExpression)
	p.registerInfix(lexer.GTE, p.parseInfixExpression)
	p.registerInfix(lexer.DOT, p.parseDotExpression)
	p.registerInfix(lexer.LPAREN, p.parseFunctionCall)
	p.registerInfix(lexer.AND, p.parseInfixExpression)
	p.registerInfix(lexer.OR, p.parseInfixExpression)
	p.registerInfix(lexer.ASSIGN, p.parseInfixExpression)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Operator: p.curToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseExpressionStatement() *ast.ExprStmt {
	stmt := &ast.ExprStmt{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Expression: p.parseExpression(LOWEST),
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		msg := fmt.Sprintf("no prefix parse function for %s found", p.curToken.Type)
		p.AddErrorMsg(&p.curToken, msg)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(lexer.SEMICOLON) &&
		!p.peekTokenIs(lexer.RBRACE) &&
		!p.peekTokenIs(lexer.RPAREN) &&
		!p.peekTokenIs(lexer.LBRACE) &&
		precedence < p.peekPrecedence() {

		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curToken.Literal}
}

func (p *Parser) parseDotExpression(left ast.Expression) ast.Expression {
	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}

	exp := &ast.DotExpression{
		BaseNode: ast.BaseNode{Token: p.curToken},
		Left:     left,
		Right:    &ast.Identifier{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curToken.Literal},
	}

	return exp
}
