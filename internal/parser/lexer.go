// internal/parser/lexer.go
package parser

import "unicode"

type TokenType string

const (
	TokenIdentifier TokenType = "IDENTIFIER"
	TokenKeyword    TokenType = "KEYWORD"
	TokenOperator   TokenType = "OPERATOR"
	TokenLiteral    TokenType = "LITERAL"
	TokenEOF        TokenType = "EOF"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(code string) *Lexer {
	return &Lexer{
		input: []rune(code),
		pos:   0,
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	ch := l.input[l.pos]

	// Identifier / keyword
	if unicode.IsLetter(ch) {
		start := l.pos
		for l.pos < len(l.input) && (unicode.IsLetter(l.input[l.pos]) || unicode.IsDigit(l.input[l.pos])) {
			l.pos++
		}
		val := string(l.input[start:l.pos])
		return Token{Type: TokenIdentifier, Value: val}
	}

	// Number literal
	if unicode.IsDigit(ch) {
		start := l.pos
		for l.pos < len(l.input) && unicode.IsDigit(l.input[l.pos]) {
			l.pos++
		}
		return Token{Type: TokenLiteral, Value: string(l.input[start:l.pos])}
	}

	// Operator
	l.pos++
	return Token{Type: TokenOperator, Value: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos++
	}
}
