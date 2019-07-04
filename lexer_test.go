package tmpl

import (
	"testing"
)

func TestLexerTextOnly(t *testing.T) {
	data := `udhas uhdaui shdauh suidah suhd asd a as
{ |sada sd \asd asd asdadiawj daw2931231 12 3123}`

	lexer, tokens, err := NewLexer(data)

	if err != nil {
		t.Fatal(err)
	}

	if lexer == nil {
		t.Fatal("error can't initialize lexer")
	}

	var token = <-tokens

	if token.Type != TokenText {
		t.Fatalf("token should be %s but got %s", TokenText, token.Type)
	}
}

func TestLexerCommentOnly(t *testing.T) {
	data := `{# {{
 {{ test }}
hello world
 {{ test }}
#}`

	lexer, tokens, err := NewLexer(data)

	if err != nil {
		t.Fatal(err)
	}

	if lexer == nil {
		t.Fatal("error can't initialize lexer")
	}

	var token = <-tokens

	if token.Type != TokenBlockCommentOpen {
		t.Fatalf("token should be %s but got %s", TokenBlockCommentOpen, token.Type)
	}

	token = <-tokens

	if token.Type != TokenBlockComment {
		t.Fatalf("token should be %s but got %s", TokenBlockComment, token.Type)
	}

	token = <-tokens

	if token.Type != TokenBlockCommentClose {
		t.Fatalf("token should be %s but got %s", TokenBlockCommentClose, token.Type)
	}
}

func TestLexerBlock(t *testing.T) {
	data := "test {{ }}"

	lexer, tokens, err := NewLexer(data)

	if err != nil {
		t.Fatal(err)
	}

	if lexer == nil {
		t.Fatal("error can't initialize lexer")
	}

	var token = <-tokens

	if token.Type != TokenText {
		t.Fatalf("token should be %s but got %s", TokenText, token.Type)
	}

	token = <-tokens

	if token.Type != TokenBlockExprOpen {
		t.Fatalf("token should be %s but got %s", TokenBlockExprOpen, token.Type)
	}
}
