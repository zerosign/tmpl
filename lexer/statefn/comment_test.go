package statefn

import (
	"testing"
)

func TestLexerCommentOnly(t *testing.T) {
	lexer := AssertNewLexer(t, `{# {{
 {{ test }}
hello world
 {{ test }}
#}`)
	AssertTokens(t, lexer, []TokenType{
		TokenBlockCommentOpen, TokenBlockComment, TokenBlockCommentClose,
	})
}
