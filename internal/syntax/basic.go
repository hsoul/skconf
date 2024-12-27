package syntax

import (
	"fmt"
	"strconv"

	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

func (p *Parser) parseInteger() ast.Expression {
	lit := &ast.Integer{BaseNode: ast.BaseNode{Token: p.curToken}}

	value, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as int", p.curToken.Literal)
		p.AddErrorMsg(&p.curToken, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseFloat() ast.Expression {
	lit := &ast.Float{BaseNode: ast.BaseNode{Token: p.curToken}}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.AddErrorMsg(&p.curToken, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseString() ast.Expression {
	return &ast.String{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{BaseNode: ast.BaseNode{Token: p.curToken}, Value: p.curTokenIs(lexer.TRUE)}
}
