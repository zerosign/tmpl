package lexer

import (
	"github.com/zerosign/tmpl/assert"
	"github.com/zerosign/tmpl/token"
	"testing"
)

func TestLexerTextOnly(t *testing.T) {
	lexer := UnsafeNewLexer("{ A*DH*AHSAHS*HASH AOSIH AHSIOHAOI SH UADUAS")
	currentToken := assert.AssertNextToken(t, lexer)
	assert.AssertToken(t, currentToken, token.TokenText)
}

func TestLexerCommentOnly(t *testing.T) {
	lexer := UnsafeNewLexer(`{# {{
 {{ test }}
hello world
 {{ test }}
#}`)

	assert.AssertTokens(t, lexer, []token.Type{
		token.TokenBlockCommentOpen, token.TokenBlockComment, token.TokenBlockCommentClose,
	})
}

func TestLexerForStmt(t *testing.T) {
	lexer := UnsafeNewLexer(`{{ for { key, value: Value } in  }}`)

}
