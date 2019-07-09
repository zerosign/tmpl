package token

import (
	"fmt"
)

// Token: Token container struct to represents current language lexer token.
//
// It uses `[]rune` (utf8 array) for storing the texts rather than string
// to unify support for utf8 for the entire flow.
//
type Token struct {
	Type  Type
	Value []rune
}

// String : string representation of token.Token
//
//
func (t Token) String() string {
	return fmt.Sprintf("Token { %s, %s }", t.Type, string(t.Value))
}
