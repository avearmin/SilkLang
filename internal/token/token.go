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

func LookupIdentType(ident string) Type {
	keywords := map[string]Type{
		"state":  State,
		"on":     On,
		"click":  Click,
		"update": Update,
		"set":    Set,
		"add":    Add,
		"remove": Remove,
	}

	if keywordType, found := keywords[ident]; found {
		return keywordType
	}
	return Ident
}
