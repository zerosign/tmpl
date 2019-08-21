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
		return nil, NotA(token.Letter)
	}

	if l.HasNext() {
		// take (<integer> | <letter> | <symbol>)
		l.TakeWhile(unicode.IsLetter, unicode.IsDigit, token.IsSymbol)
	}

	l.Emit(token.Ident)

	return nil, nil
}
