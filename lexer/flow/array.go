package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
)

//
// []
// ["test"]
// [data]
// ["test", 1]
//
func LexArray(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter array expr")
	defer log.Debug().Msg("exit array expr")

	l.CursorMut().Next()
	l.Emit(token.TokenBracketOpen)

	// TODO: @zerosign

	l.CursorMut().Next()
	l.Emit(token.TokenBracketClose)

	return nil, nil
}
