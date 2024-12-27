package lexer

import "fmt"

type Position struct {
	Line   int // 当前行号
	Column int // 当前列号
}

type Lexer struct {
	input        string
	position     int      // 当前字符的位置
	readPosition int      // 当前读取位置（在当前字符之后）
	ch           byte     // 当前正在查看的字符
	pos          Position // 当前解析位置
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		pos: Position{
			Line:   1,
			Column: 0,
		},
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.pos.Line++
		l.pos.Column = 0
	} else {
		l.pos.Column++
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: ASSIGN, Literal: string(l.ch)}
		}
	case '+':
		tok = Token{Type: PLUS, Literal: string(l.ch)}
	case '-':
		if l.peekChar() == '-' { // 注释
			l.readChar() // 跳过第二个'-'
			return l.readComment()
		}
		tok = Token{Type: MINUS, Literal: string(l.ch)}
	case '*':
		tok = Token{Type: MULTIPLY, Literal: string(l.ch)}
	case '/':
		tok = Token{Type: DIVIDE, Literal: string(l.ch)}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: LT, Literal: string(l.ch)}
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: GT, Literal: string(l.ch)}
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch)}
		}
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case '.':
		tok = Token{Type: DOT, Literal: string(l.ch)}
	case ',':
		tok = Token{Type: COMMA, Literal: string(l.ch)}
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch)}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch)}
	case '{':
		tok = Token{Type: LBRACE, Literal: string(l.ch)}
	case '}':
		tok = Token{Type: RBRACE, Literal: string(l.ch)}
	case '[':
		tok = Token{Type: LBRACKET, Literal: string(l.ch)}
	case ']':
		tok = Token{Type: RBRACKET, Literal: string(l.ch)}
	case ';':
		tok = Token{Type: SEMICOLON, Literal: string(l.ch)}
	case 0:
		tok = Token{Type: EOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			tok.Pos = l.pos
			fmt.Println(tok)
			return tok
		} else if isDigit(l.ch) {
			tok = l.readNumber()
			fmt.Println(tok)
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch)}
		}
	}

	l.readChar()
	tok.Pos = l.pos
	fmt.Println(tok)
	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() Token {
	startPosition := l.position
	isFloat := false

	for isDigit(l.ch) {
		l.readChar()
	}

	if l.ch == '.' && isDigit(l.peekChar()) {
		isFloat = true
		l.readChar() // consume the dot
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	if isFloat {
		return Token{
			Type:    FLOAT,
			Literal: l.input[startPosition:l.position],
			Pos:     l.pos,
		}
	}

	return Token{
		Type:    INTEGER,
		Literal: l.input[startPosition:l.position],
		Pos:     l.pos,
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' || l.ch == '-' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readComment() Token {
	position := l.position + 1
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return Token{
		Type:    COMMENT,
		Literal: l.input[position:l.position],
		Pos:     l.pos,
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
