package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes/block"
	rutil "github.com/zerosign/tmpl/runes/util"
)

//
// <template> : (<expression> | <assignment> | <comment> | <text>) <template>
//
// template are also terminal state.
//
func LexTemplate(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter LexTemplate")
	defer log.Debug().Msg("exit LexTemplate")

	// terminal
	if !l.Available() {
		return nil, nil
	}

	value := l.RunesAhead()

	// test if it's expression
	// then assignment
	// then comment
	// then text
	if rutil.HasPrefix(value, block.OpenExpr) {
		return LexBlockExprOpen, nil
	} else if rutil.HasPrefix(value, block.OpenAssign) {
		return LexBlockAssignOpen, nil
	} else if rutil.HasPrefix(value, block.OpenComment) {
		return LexBlockCommentOpen, nil
	} else {
		// anything other than above
		return LexText, nil
	}
}
