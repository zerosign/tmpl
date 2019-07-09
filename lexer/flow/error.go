package flow

import (
	"fmt"
	"github.com/zerosign/tmpl/runes"
)

const (
	Separator rune = ','
)

// NoMatchToken: Error if there were no matching token
//
func NoMatchToken(scope string, expected [][]rune) error {
	return fmt.Errorf("no match token for scope %s, expected: (%s)", scope, runes.Join(expected, Separator))
}

// NotA: return an error that give whether current token is not
//
func NotA(scope string) error {
	return fmt.Errorf("current token is not %s", scope)
}
