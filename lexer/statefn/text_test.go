package statefn

import (
	"testing"
)

func TestLexerTextOnly(t *testing.T) {
	lexer := AssertNewLexer(t, "{ A*DH*AHSAHS*HASH AOSIH AHSIOHAOI SH UADUAS")
	token := AssertNextToken(t, lexer)
	AssertToken(t, token, TokenText)
}
