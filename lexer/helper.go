package lexer

import (
	"fmt"
	"github.com/zerosign/tmpl/base"
)

// AssertNewLexer: Utility function for creating new lexer
//
// will raise fatal error if fails
//
func UnsafeNewLexer(input string) base.Lexer {
	nlexer, err := DefaultLexer(input)

	fmt.Print(input)

	if err != nil {
		panic(err)
	}

	return nlexer
}
