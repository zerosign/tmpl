package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"log"
)

//
// <template> : (<expression> | <assignment> | <comment> | <text>) <template>
//
// template are also terminal state.
//
func LexTemplate(l base.Lexer) (Flow, error) {
	log.Println("enter LexTemplate")
	defer log.Println("exit LexTemplate")

	// terminal
	if !l.Available() {
		return nil, nil
	}

	value := l.RunesAhead()

	// test if it's expression
	// then assignment
	// then comment
	// then text
	if runes.HasPrefix(value, token.BlockExprOpen) {
		return LexBlockExprOpen, nil
	} else if runes.HasPrefix(value, token.BlockAssignOpen) {
		return LexBlockAssignOpen, nil
	} else if runes.HasPrefix(value, token.BlockCommentOpen) {
		return LexBlockCommentOpen, nil
	} else {
		// anything other than above
		return LexText, nil
	}
}
