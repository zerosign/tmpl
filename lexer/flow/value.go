package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
)

func LexValue(l base.Lexer) (Flow, error) {
	ch := l.CurrentRune()

	// check for primitive type
	// this include quoted string or digit
	if token.IsPrimitive(ch) {
		return LexPrimitive, nil
	} else if token.IsBraceOpen(ch) {
		// map type
		return LexMap, nil
	} else if token.IsBracketOpen(ch) {
		// array type
		return LexArray, nil
	} else {
		// if value is not primitive | map | array
		return nil, UnsupportedValueType()
	}
}
