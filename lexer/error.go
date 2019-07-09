package lexer

import (
	"fmt"
)

// InvalidUtfInput: return an error for invalid utf8 input
//
func InvalidUtfInput() error {
	return fmt.Errorf("invalid utf8 input")
}
