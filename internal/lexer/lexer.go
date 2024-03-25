package lexer

import "github.com/avearmin/SilkLang/internal/token"

type Lexer struct {
	input   string
	pos     int  // current pos in input (current char)
	readPos int  // current reading pos in input (after current char)
	char    byte // current char being looked at
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Get the next char and advance our pos in the input string.
// Note that a 0 byte represents an ASCII NUL char, and in
// this context that means an EOF.
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespaces()

	switch l.char {
	case '=':
		tok = token.New(token.Assign, l.char)
	case '+':
		tok = token.New(token.Plus, l.char)
	case ':':
		tok = token.New(token.Colon, l.char)
	case ';':
		tok = token.New(token.SemiColon, l.char)
	case '{':
		tok = token.New(token.LBrace, l.char)
	case '}':
		tok = token.New(token.RBrace, l.char)
	case '#':
		tok = token.New(token.IdSelector, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.Eof
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupIdentType(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.Int
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.New(token.Illegal, l.char)
			return tok
		}
	}

	l.readChar()
	return tok
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_') || (char == '-')
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) consumeWhitespaces() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdent() string {
	pos := l.pos
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}
