package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
)

// LexBlockAssignOpen :
//
func LexBlockAssignOpen(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter block assign open")
	defer log.Debug().Msg("exit block assign open")

	l.CursorMut().Incr(len(token.BlockAssignOpen))
	l.Emit(token.TokenBlockAssignOpen)
	return LexBlockAssign, nil
}

//
// block region contains :
// - expression statement
//
func LexBlockAssign(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter block assign")
	defer log.Debug().Msg("exit block assign")

	l.Ignore(token.IsWhitespace)

	return LexBlockAssignClose, nil
}

func LexBlockAssignClose(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter block assign close")
	defer log.Debug().Msg("exit block assign close")

	l.CursorMut().Incr(len(token.BlockAssignClose))
	l.Emit(token.TokenBlockAssignClose)
	return LexText, nil
}
