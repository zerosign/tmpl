package lexer

import (
	"fmt"
)

const (
	Separator rune = ','
)

// NoMatchToken: Error if there were no matching token
//
func NoMatchToken(scope string, expected [][]rune) error {
	return fmt.Errorf("no match token for scope %s, expected: (%s)", scope, RuneStrSep(expected, Separator))
}

// NotA: return an error that give whether current token is not
//
func NotA(scope string) error {
	return fmt.Errorf("current token is not %s", scope)
}

// InvalidUtfInput: return an error for invalid utf8 input
//
func InvalidUtfInput() error {
	return fmt.Errorf("invalid utf8 input")
}
