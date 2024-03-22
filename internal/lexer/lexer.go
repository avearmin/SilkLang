package lexer

type Lexer struct {
	input   string
	pos     int  // current pos in input (current char)
	readPos int  // current reading pos in input (after current char)
	char    byte // current char being looked at
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
