package syntax

import (
	"fmt"

	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

func (p *Parser) parseTable() ast.Expression {
	table := &ast.TableDef{
		BaseNode: ast.BaseNode{
			Token: p.curToken,
		},
		Properties: []*ast.PropertyDef{},
	}

	if p.peekTokenIs(lexer.RBRACE) {
		p.nextToken()
		return table
	}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) {
		if p.curTokenIs(lexer.COMMENT) {
			p.nextToken()
			continue
		}

		if p.curTokenIs(lexer.LBRACKET) || (p.curTokenIs(lexer.IDENTIFIER) && !p.peekTokenIs(lexer.DOT)) {
			var key ast.Expression
			if p.curTokenIs(lexer.LBRACKET) { // Handle array-style indexing with [key]
				p.nextToken() // consume '['
				key = p.parseExpression(LOWEST)
				if key == nil {
					return nil
				}
				if !p.expectPeek(lexer.RBRACKET) {
					return nil
				}
			} else {
				if !p.curTokenIs(lexer.IDENTIFIER) { // Handle regular identifiers
					msg := fmt.Sprintf("expected identifier or [key] as table key, got %s", p.curToken.Type)
					p.AddErrorMsg(&p.curToken, msg)
					return nil
				}
				key = &ast.Identifier{
					BaseNode: ast.BaseNode{Token: p.curToken},
					Value:    p.curToken.Literal,
				}
			}

			if !p.expectPeek(lexer.ASSIGN) {
				return nil
			}
			p.nextToken() // consume '='

			value := p.parseExpression(LOWEST)
			if value == nil {
				return nil
			}

			table.Properties = append(table.Properties, &ast.PropertyDef{
				BaseNode: ast.BaseNode{Token: p.curToken},
				Key:      key,
				Value:    value,
			})
		} else { // array
			value := p.parseExpression(LOWEST)
			if value == nil {
				return nil
			}

			table.Properties = append(table.Properties, &ast.PropertyDef{
				BaseNode: ast.BaseNode{Token: p.curToken},
				Key:      nil,
				Value:    value,
			})
		}

		if p.peekTokenIs(lexer.COMMA) { // Handle comma separator ;
			p.nextToken()
			if p.peekTokenIs(lexer.RBRACE) {
				break
			}
			p.nextToken()
		} else if p.peekTokenIs(lexer.RBRACE) {
			break
		} else if !p.peekTokenIs(lexer.RBRACE) {
			msg := fmt.Sprintf("expected next token to be COMMA or RBRACE, got %s instead", p.peekToken.Type)
			p.AddErrorMsg(&p.curToken, msg)
			return nil
		}
	}

	if !p.expectPeek(lexer.RBRACE) {
		return nil
	}

	return table
}
