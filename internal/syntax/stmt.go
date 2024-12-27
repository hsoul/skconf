package syntax

import (
	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.SKILL:
		return p.parseSkillDefinition()
	case lexer.STATE:
		return p.parseStateDefinition()
	case lexer.IF:
		return p.parseIfStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	case lexer.COMMENT:
		return p.parseComment()
	case lexer.VAR:
		return p.parseVarStatement()
	case lexer.FOR:
		return p.parseForStatement()
	case lexer.BREAK:
		return &ast.BreakStatement{BaseNode: ast.BaseNode{Token: p.curToken}}
	case lexer.CONTINUE:
		return &ast.ContinueStatement{BaseNode: ast.BaseNode{Token: p.curToken}}
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Imports:    []ast.ImportStatement{},
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != lexer.EOF && len(p.errors) == 0 {
		if p.curToken.Type == lexer.IMPORT {
			if stmt := p.parseImportStatement(); stmt != nil {
				program.Imports = append(program.Imports, *stmt)
			}
		} else if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	if len(p.errors) > 0 {
		return nil
	}

	return program
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{
		BaseNode: ast.BaseNode{Token: p.curToken},
	}

	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		BaseNode: ast.BaseNode{Token: p.curToken},
		Value:    p.curToken.Literal,
	}

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	p.nextToken() // skip =

	stmt.Value = p.parseExpression(LOWEST)
	if stmt.Value == nil {
		return nil
	}

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{
		BaseNode: ast.BaseNode{Token: p.curToken},
	}

	p.nextToken() // consume 'for'

	if p.curTokenIs(lexer.IDENTIFIER) { // Check if this is a range-based for loop
		if p.peekTokenIs(lexer.COMMA) {
			stmt.IsRangeForm = true // for key, value = range expr { }
			stmt.Key = &ast.Identifier{
				BaseNode: ast.BaseNode{Token: p.curToken},
				Value:    p.curToken.Literal,
			}

			p.nextToken() // consume key
			p.nextToken() // consume comma

			if !p.curTokenIs(lexer.IDENTIFIER) {
				p.AddErrorMsg(&p.curToken, "expected identifier after comma in range statement")
				return nil
			}

			stmt.Value = &ast.Identifier{
				BaseNode: ast.BaseNode{Token: p.curToken},
				Value:    p.curToken.Literal,
			}

			if !p.expectPeek(lexer.ASSIGN) {
				return nil
			}

			if !p.expectPeek(lexer.RANGE) {
				return nil
			}

			p.nextToken() // consume range
			stmt.RangeValue = p.parseExpression(LOWEST)

			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}

			stmt.Body = p.parseBlockStatement()
			return stmt
		}
	}

	if !p.curTokenIs(lexer.VAR) && !p.curTokenIs(lexer.SEMICOLON) { // Check if this is a condition-only for loop
		stmt.Condition = p.parseExpression(LOWEST) // for condition { }

		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}

		stmt.Body = p.parseBlockStatement()
		return stmt
	}

	if !p.curTokenIs(lexer.SEMICOLON) { // Classic for loop: for init; condition; post { }
		stmt.Init = p.parseStatement()
	}

	if !p.curTokenIs(lexer.SEMICOLON) {
		if !p.expectPeek(lexer.SEMICOLON) {
			return nil
		}
	} else {
		p.nextToken() // skip semicolon
	}

	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.curTokenIs(lexer.SEMICOLON) {
		if !p.expectPeek(lexer.SEMICOLON) {
			return nil
		}
	} else {
		p.nextToken() // skip semicolon
	}

	if !p.curTokenIs(lexer.LBRACE) {
		p.nextToken()
		stmt.Post = &ast.ExprStmt{
			BaseNode:   ast.BaseNode{Token: p.curToken},
			Expression: p.parseExpression(LOWEST),
		}
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()
	return stmt
}

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
	}

	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}

	stmt.Value = p.parseExpression(LOWEST)
	return stmt
}

func (p *Parser) parseIfStatement() ast.Expression {
	if !p.curTokenIs(lexer.IF) {
		return nil
	}

	stmt := &ast.IfStatement{
		BaseNode:     ast.BaseNode{Token: p.curToken},
		Alternatives: []*ast.ElseStatement{},
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	for p.peekTokenIs(lexer.ELSE) { // Parse else-if and else chains
		p.nextToken() // consume 'else'

		if p.peekTokenIs(lexer.IF) {
			p.nextToken() // consume 'if'
			p.nextToken() // move to condition

			elseStmt := &ast.ElseStatement{
				BaseNode:  ast.BaseNode{Token: p.curToken},
				Condition: p.parseExpression(LOWEST),
			}

			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}

			elseStmt.Consequence = p.parseBlockStatement()
			stmt.Alternatives = append(stmt.Alternatives, elseStmt)
		} else {
			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}

			elseStmt := &ast.ElseStatement{
				BaseNode:    ast.BaseNode{Token: p.curToken},
				Condition:   nil,
				Consequence: p.parseBlockStatement(),
			}
			stmt.Alternatives = append(stmt.Alternatives, elseStmt)
			break // after 'else', no more branches are possible
		}
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
	}

	p.nextToken()

	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.ReturnValue = p.parseExpression(LOWEST)
	}

	return stmt
}

func (p *Parser) parseComment() *ast.CommentStatement {
	stmt := &ast.CommentStatement{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Value: p.curToken.Literal,
	}
	p.nextToken() // Move past the comment
	return stmt
}

func (p *Parser) parseBlockStatement() *ast.CodeBlock {
	block := &ast.CodeBlock{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Statements: []ast.Statement{},
	}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.COMMENT) {
			p.nextToken()
			continue
		}
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}
