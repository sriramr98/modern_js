package lexer

import (
	"github.com/sriramr98/modern_js/token"
)

// to support full unicode range and emoji's, ch should be a Rune and all parsing logic needs to be changed
type Lexer struct {
	input        string
	position     int  // current position in the input
	readPosition int  // current reading position in input
	ch           byte // current chat under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCharAndMoveToNext() // to make sure struct variables are initialized to valid values respective to the input
	return l
}

func (l *Lexer) readCharAndMoveToNext() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // 0 is the ascii code for NUL
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition < len(l.input) {
		return l.input[l.readPosition]
	}

	return 0
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readCharAndMoveToNext()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	// Operators
	case '=':
		if l.peekChar() == '=' {
			// Case for ==
			tok = l.createTwoCharToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '%':
		tok = newToken(token.MOD, l.ch)
	case '!':
		if l.peekChar() == '=' {
			// Case for !=
			tok = l.createTwoCharToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	// Conditionals
	case '<':
		if l.peekChar() == '=' {
			// Case for <=
			tok = l.createTwoCharToken(token.LT_OR_EQ)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			// Case for >=
			tok = l.createTwoCharToken(token.GT_OR_EQ)
		} else {
			tok = newToken(token.GT, l.ch)
		}
		// Delimiters
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		// Detects Keywords
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// Detects Numbers
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readCharAndMoveToNext()

	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readCharAndMoveToNext()
	}

	return l.input[pos:l.position]
}

// This implementation is super simplified and doesn't support floats, hexadecimal etc..
func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readCharAndMoveToNext()
	}

	return l.input[pos:l.position]
}

// Creates a token with the current char and the next char as literal
func (l *Lexer) createTwoCharToken(tok token.TokenType) token.Token {
	ch := l.ch
	l.readCharAndMoveToNext()
	return token.Token{Type: tok, Literal: string(ch) + string(l.ch)}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_' || ch == '-'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func newToken(tokType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokType, Literal: string(ch)}
}
