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

	TagSelector   = "TAG_SELECTOR"
	IdSelector    = "ID_SELECTOR"
	ClassSelector = "CLASS_SELECTOR"

	Ident = "IDENT"
)

type Token struct {
	Type    Type
	Literal string
}
