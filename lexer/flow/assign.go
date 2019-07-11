package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"log"
)

// LexBlockAssignOpen :
//
func LexBlockAssignOpen(l base.Lexer) (Flow, error) {
	log.Println("enter block assign")
	l.CursorMut().Incr(len(token.BlockAssignOpen))
	l.Emit(token.TokenBlockAssignOpen)
	return LexBlockAssign, nil
}

//
// block region contains :
// - expression statement
//
func LexBlockAssign(l base.Lexer) (Flow, error) {

	return nil, nil
}

func LexBlockAssignClose(l base.Lexer) (Flow, error) {
	log.Println("exit block assign")
	l.CursorMut().Incr(len(token.BlockAssignClose))
	l.Emit(token.TokenBlockAssignClose)
	return LexText, nil
}
