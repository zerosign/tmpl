package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"unicode"
)

// LexSignedNumber: Use signed number lexer by default
//
//
func LexSignedNumber(l base.Lexer) (Flow, error) {

	currentRune := l.CurrentRune()

	if currentRune == runes.Negative ||
		currentRune == runes.Positive {
		l.Emit(token.Sign)

		if l.Available() {
			l.Advance()
		}
	}

	currentRune = l.CurrentRune()

	if unicode.IsDigit(currentRune) {
		return LexNumber, nil
	} else {
		return nil, NotA(token.Digit)
	}
}

// LexNumber : function to lex all numeric value
//
// This identify :
// - integer
// - double (if in the middle of lexing founds '.')
//
func LexNumber(l base.Lexer) (Flow, error) {

	// handle (<digit_without_zero>+ <digit>*) | <digit_zero>
	if l.CurrentRune() == '0' {
		l.Advance()
		ch := l.CurrentRune()

		// no runes means the value is a integer ~ 0
		if int(ch) == 0 { // integer case
			l.Emit(token.Integer)
			return nil, nil
		} else if token.IsDot(ch) { // double case
			return LexFraction, nil
		} else {
			return nil, ZeroPaddedInteger()
		}
	}

	for l.Available() {
		ch := l.CurrentRune()

		if unicode.IsDigit(ch) {
			l.Advance()
		} else if token.IsDot(ch) {
			return LexFraction, nil
		}
	}

	l.Emit(token.Integer)

	return nil, nil
}

func LexFraction(l base.Lexer) (Flow, error) {

	for l.Available() {
		currentRune := l.CurrentRune()

		if unicode.IsDigit(currentRune) {
			l.Advance()
		} else if token.IsWhitespace(currentRune) {
			break
		} else {
			return nil, NotA(token.Digit)
		}
	}

	l.Emit(token.Double)

	return nil, nil
}

func LexString(l base.Lexer) (Flow, error) {
	var lastRune rune

	// first token '"', if not quote token then
	// return error
	if l.CurrentRune() != runes.Quote {
		return nil, NotA(token.Quote)
	}

	// consume all string until token.Quote, unless
	// last rune are token.SymbolEscape
	// stop invariant
	// - ends with token.Quote unless last token is token.SymbolEscape
	//
	for {
		currentRune := l.CurrentRune()

		if currentRune == runes.Quote {
			if lastRune != runes.Escape {
				break
			}
		}

		lastRune = currentRune

		if l.Available() {
			l.Advance()
		} else {
			return nil, NotA(token.Quote)
		}
	}

	l.Emit(token.String)

	return nil, nil
}

func LexPrimitive(l base.Lexer) (Flow, error) {

	return nil, nil
}
