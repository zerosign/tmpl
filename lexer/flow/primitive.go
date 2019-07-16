package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"unicode"
)

func LexInteger(l base.Lexer) (Flow, error) {

	// no zero padded integer at beginning
	// 0x

	if l.CurrentRune() == '0' {
		l.Next()
		if unicode.IsDigit(l.CurrentRune()) {
			return nil, ZeroPaddedInteger()
		}
	}

	// take characters while it's a digit
	l.TakeWhile(unicode.IsDigit)
	l.Emit(token.TokenInteger)

	return nil, nil
}

func LexString(l base.Lexer) (Flow, error) {
	if l.CurrentRune() != token.Quote {
		return nil, NotA(token.TokenQuote)
	} else {
		l.Next()
	}

	// TODO: stop invariant
	for l.CurrentRune() != token.Quote {
		l.Next()
	}

	return nil, nil
}

func LexDouble(l base.Lexer) (Flow, error) {

	return nil, nil
}

func LexPrimitive(l base.Lexer) (Flow, error) {
	return nil, nil
}
