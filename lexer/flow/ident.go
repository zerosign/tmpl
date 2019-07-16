package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"unicode"
)

//
// <letter> (<integer> | <letter> | <symbol>)*
//
func LexIdent(l base.Lexer) (Flow, error) {
	//
	// should be a letter
	if !unicode.IsLetter(l.CurrentRune()) {
		return nil, NotA(token.TokenLetter)
	}

	if l.HasNext() {
		// take (<integer> | <letter> | <symbol>)
		l.TakeWhile(unicode.IsLetter, unicode.IsDigit, token.IsSymbol)
	}

	l.Emit(token.TokenIdent)

	return nil, nil
}
