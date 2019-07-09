package statefn

import (
	"github.com/zerosign/tmpl/lexer/token"
	"testing"
)

func TestLexerBlock(t *testing.T) {
	lexer := AssertNewLexer(t, "test {{ }}")

	AssertTokens(t, lexer, []token.TokenType{
		TokenText, TokenBlockExprOpen,
	})
}
