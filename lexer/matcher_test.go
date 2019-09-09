package lexer

import (
	"testing"
)

func TestIsNewline(t *testing.T) {
	tests := []rune{'\n', 'r'}

	for _, test := range tests {
		if !IsNewline(test) {
			t.Errorf("'%v' should be read as newline", test)
		}
	}
}

func TestIsWhitespaceOnly(t *testing.T) {
	tests = []rune{' ', '\t'}

	for _, test := range tests {
		if !IsWhitespaceOnly(test) {
			t.Errorf("'%v' should be read as whitespace only", test)
		}
	}
}
