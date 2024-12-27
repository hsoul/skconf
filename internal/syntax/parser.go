package syntax

import (
	"fmt"
	"runtime"

	"github.com/hsoul/skconf/internal/ast"
	"github.com/hsoul/skconf/internal/lexer"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	curToken       lexer.Token
	peekToken      lexer.Token
	errors         []string
	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn
	fileName       string
}

func New(l *lexer.Lexer, fileName string) *Parser {
	p := &Parser{
		l:              l,
		errors:         []string{},
		prefixParseFns: make(map[lexer.TokenType]prefixParseFn),
		infixParseFns:  make(map[lexer.TokenType]infixParseFn),
		fileName:       fileName,
	}

	p.registerPrecedenceFuncs()

	p.nextToken() // Read two tokens, so curToken and peekToken are both set
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

	for p.curTokenIs(lexer.COMMENT) {
		p.curToken = p.peekToken
		p.peekToken = p.l.NextToken()
	}

	for p.peekTokenIs(lexer.COMMENT) {
		p.peekToken = p.l.NextToken()
	}
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool { // Check if the next token is t and consume it
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	msg := fmt.Sprintf("unexpected token %q, expected %q", p.peekToken.Literal, t.TokenLiteral())
	p.AddErrorMsg(&p.peekToken, msg)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) AddErrorMsg(tok *lexer.Token, msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	if tok != nil {
		msg = fmt.Sprintf("%s:%d: %s:%d:%d: %s ", file, line, p.fileName, tok.Pos.Line, tok.Pos.Column, msg)
	} else {
		msg = fmt.Sprintf("%s:%d: %s", file, line, msg)
	}
	p.errors = append(p.errors, msg)
	fmt.Println(msg)
}
