package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
)

// LexArray: lexing array declaration
// []
// ["test"]
// [data]
// ["test", 1]
//
// currently we only support for 1D array.
//
func LexArray(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter array expr")
	defer log.Debug().Msg("exit array expr")

	l.CursorMut().Next()
	l.Emit(token.OpenBracket)

	// TODO: @zerosign

	l.CursorMut().Next()
	l.Emit(token.CloseBracket)

	return nil, nil
}
