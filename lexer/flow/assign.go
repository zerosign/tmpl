package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"log"
)

// LexBlockAssignOpen :
//
func LexBlockAssignOpen(l base.Lexer) (Flow, error) {
	log.Println("enter block assign open")
	defer log.Println("exit block assign open")

	l.CursorMut().Incr(len(token.BlockAssignOpen))
	l.Emit(token.TokenBlockAssignOpen)
	return LexBlockAssign, nil
}

//
// block region contains :
// - expression statement
//
func LexBlockAssign(l base.Lexer) (Flow, error) {

	l.Ignore(token.IsWhitespace)

	return LexBlockAssignClose, nil
}

func LexBlockAssignClose(l base.Lexer) (Flow, error) {
	log.Println("enter block assign close")
	defer log.Println("exit block assign close")

	l.CursorMut().Incr(len(token.BlockAssignClose))
	l.Emit(token.TokenBlockAssignClose)
	return LexText, nil
}
