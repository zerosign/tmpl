package flow

import (
	"fmt"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

type textCase int

const (
	upperCase textCase = iota
	lowerCase
)

func (t textCase) String() {
	if t == upperCase {
		return "<uppercase>"
	} else {
		return "<lowercase>"
	}
}

const (
	Separator rune = ','
)

// NoMatchToken: Error if there were no matching token
//
func NoMatchToken(scope string, expected [][]rune) error {
	return fmt.Errorf("no match token for scope %s, expected: (%s)", scope, runes.Join(expected, Separator))
}

func ZeroPaddedInteger() error {
	return fmt.Errorf("integer `%s` value shouldn't be padded with 0")
}

func

// NotA: return an error that give whether current token is not
//
func NotA(tt token.Type) error {
	return fmt.Errorf("current token is not %s", tt)
}

func CaseError(expected, result textCase) error {
	return fmt.Errorf("expected case: %s got %s", expected, result)
}
