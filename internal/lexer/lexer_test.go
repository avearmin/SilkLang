package lexer

import (
	"github.com/avearmin/SilkLang/internal/token"
	"testing"
)

func TestParse(t *testing.T) {
	input := `
state clicks = 0;

on click #my-button {
  clicks = clicks + 1;
  update #display {
    set content: clicks;
  }
}
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.State, "state"},
		{token.Ident, "clicks"},
		{token.Assign, "="},
		{token.Int, "0"},
		{token.SemiColon, ";"},
		{token.On, "on"},
		{token.Click, "click"},
		{token.IdSelector, "#"},
		{token.Ident, "my-button"},
		{token.LBrace, "{"},
		{token.Ident, "clicks"},
		{token.Assign, "="},
		{token.Ident, "clicks"},
		{token.Plus, "+"},
		{token.Int, "1"},
		{token.SemiColon, ";"},
		{token.Update, "update"},
		{token.IdSelector, "#"},
		{token.Ident, "display"},
		{token.LBrace, "{"},
		{token.Set, "set"},
		{token.Ident, "content"},
		{token.Colon, ":"},
		{token.Ident, "clicks"},
		{token.SemiColon, ";"},
		{token.RBrace, "}"},
		{token.RBrace, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
