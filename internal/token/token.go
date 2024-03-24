package token

type Type string

const (
	Illegal = "ILLEGAL"
	Eof     = "EOF"

	Int = "INT"

	Assign = "="
	Plus   = "+"

	Colon     = ":"
	SemiColon = ";"
	LBrace    = "{"
	RBrace    = "}"
	LParen    = "("
	RParen    = ")"

	State  = "STATE"
	On     = "ON"
	Click  = "CLICK"
	Update = "UPDATE"
	Set    = "SET"
	Add    = "ADD"
	Remove = "REMOVE"

	IdSelector    = "#"
	ClassSelector = "."

	Ident = "IDENT"
)

type Token struct {
	Type    Type
	Literal string
}

func New(tokenType Type, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}
