package flow

import (
	"fmt"
	rutil "github.com/zerosign/tmpl/runes/util"
	"github.com/zerosign/tmpl/token"
)

type textCase int

const (
	upperCase textCase = iota
	lowerCase
)

func (t textCase) String() string {
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
	return fmt.Errorf("no match token for scope %s, expected: (%s)", scope, rutil.Join(expected, Separator))
}

func ZeroPaddedInteger() error {
	return fmt.Errorf("integer value shouldn't be padded with 0")
}

// NotA: return an error that give whether current token is not
//
func NotA(tt token.Type) error {
	return fmt.Errorf("current token is not %s", tt)
}

func CaseError(expected, result textCase) error {
	return fmt.Errorf("expected case: %s got %s", expected, result)
}

func UnsupportedValueType() error {
	return fmt.Errorf("unsupported value type")
}
