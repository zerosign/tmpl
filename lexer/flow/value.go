package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
)

func LexValue(l base.Lexer) (Flow, error) {
	ch := l.CurrentRune()

	// check for primitive type
	if token.IsPrimitive(ch) {
		return LexPrimitive, nil
	} else if token.IsBraceOpen(ch) {
		// map type
		return LexMap, nil
	} else if token.IsParaOpen(ch) {
		// array type
		return LexArray, nil
	} else {
		// TODO : add new error
		return nil, nil
	}
}
