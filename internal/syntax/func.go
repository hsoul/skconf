package syntax

import (
	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseFunction() ast.Expression {
	funcLit := &ast.FunctionDef{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
	}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	funcLit.Parameters = p.parseFunctionParameters()
	if funcLit.Parameters == nil {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	funcLit.Body = p.parseBlockStatement()
	if funcLit.Body == nil {
		return nil
	}

	return funcLit
}

func (p *Parser) parseFunctionCall(function ast.Expression) ast.Expression {
	exp := &ast.FunctionCall{
		BaseNode:  ast.BaseNode{Token: p.curToken},
		Function:  function,
		Arguments: p.parseExpressionList(lexer.RPAREN),
	}
	return exp
}
