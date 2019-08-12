package token

import (
	"testing"
	"github.com/zerosign/tmpl/runes"
)

func TestIsRuneEq(t *testing.T) {
	matcher := IsRuneEq(runes.BraceOpen)

	if !matcher(runes.BraceOpen) {
		t.Errorf("should be able to match '%v'", runes.BraceOpen)
	}
}

func TestIsWhitespace(t *testing.T) {
	if !IsWhitespace(' ') {
		t.Errorf("should be able to match '%v'", " ")
	}
}

func TestIsSymbol(t *testing.T) {

}
