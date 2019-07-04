package tmpl

import (
	"fmt"
)

const (
	Separator rune = ','
)

// Error if there were no matching token
//
func NoMatchToken(scope string, expected [][]rune) error {
	return fmt.Errorf("no match token for scope %s, expected: (%s)", scope, RuneStrSep(expected, Separator))
}

func NotA(scope string) error {
	return fmt.Errorf("current token is not %s", scope)
}

func InvalidUtfInput() error {
	return fmt.Errorf("invalid utf8 input")
}
