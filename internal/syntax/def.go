package syntax

import (
	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

func (p *Parser) parseSkillDefinition() *ast.SkillDef {
	skill := &ast.SkillDef{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Properties: []*ast.PropertyDef{},
	}

	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}

	skill.Name = &ast.Identifier{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.COMMENT) {
			p.nextToken()
			continue
		}

		if p.curTokenIs(lexer.IDENTIFIER) {
			name := p.curToken.Literal
			identifier := p.curToken
			if p.peekTokenIs(lexer.ASSIGN) {
				p.nextToken() // skip the "="
				p.nextToken()

				if !p.curTokenIs(lexer.FUNC) {
					expr := p.parseExpression(LOWEST)
					if expr != nil {
						skill.Properties = append(skill.Properties, &ast.PropertyDef{
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Token: identifier,
								},
								Value: name,
							},
							Value: expr,
						})
					}
				} else {
					method := p.parseFunction()
					if _, ok := method.(*ast.FunctionDef); !ok {
						return nil
					}
					if method != nil {
						skill.Properties = append(skill.Properties, &ast.PropertyDef{
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Token: identifier,
								},
								Value: name,
							},
							Value: method,
						})
					}
				}
			}
		}

		if !p.expectPeek(lexer.COMMA) && !p.curTokenIs(lexer.RBRACE) {
			return nil
		}

		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
		}

		p.nextToken()
	}

	return skill
}

func (p *Parser) parseStateDefinition() *ast.StateDef {
	state := &ast.StateDef{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Properties: []*ast.PropertyDef{},
	}

	if !p.expectPeek(lexer.IDENTIFIER) {
		return nil
	}

	state.Name = &ast.Identifier{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curToken.Type == lexer.IDENTIFIER {
			name := p.curToken.Literal
			identifier := p.curToken

			if p.peekTokenIs(lexer.ASSIGN) {
				p.nextToken() // 跳过=
				p.nextToken()

				if p.peekTokenIs(lexer.LBRACE) {
					if table := p.parseTable(); table != nil {
						state.Properties = append(state.Properties, &ast.PropertyDef{
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Token: identifier,
								},
								Value: name,
							},
							Value: table,
						})
					}
				} else if !p.curTokenIs(lexer.FUNC) {
					expr := p.parseExpression(LOWEST)
					if expr != nil {
						state.Properties = append(state.Properties, &ast.PropertyDef{
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Token: identifier,
								},
								Value: name,
							},
							Value: expr,
						})
					}
				} else {
					method := p.parseFunction()
					if _, ok := method.(*ast.FunctionDef); !ok {
						return nil
					}
					if method != nil {
						state.Properties = append(state.Properties, &ast.PropertyDef{
							Key: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Token: identifier,
								},
								Value: name,
							},
							Value: method,
						})
					}
				}
			}
		}

		if !p.expectPeek(lexer.COMMA) && !p.curTokenIs(lexer.RBRACE) {
			return nil
		}
		p.nextToken()
	}

	return state
}
