package lexer

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/lexer/flow"
)

// UnsafeDefaultLexer: Utility function for creating new lexer with default flow.
//
// will raise fatal error if fails
//
func UnsafeDefaultLexer(input string) base.Lexer {
	nlexer, err := DefaultLexer(input)

	if err != nil {
		panic(err)
	}

	return nlexer
}

// UnsafeNewLexer: Utility function for creating new lexer with given flow.
//
func UnsafeNewLexer(input string, f flow.Flow) base.Lexer {
	nlexer, err := NewLexer(input, f)

	if err != nil {
		panic(err)
	}

	return nlexer
}
