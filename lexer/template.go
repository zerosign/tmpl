package lexer

import (
	"github.com/rs/zerolog/log"
)

//
// <template> : (<expression> | <assignment> | <comment> | <text>) <template>
//
// template are also terminal state.
//
func TemplateFlow(l Lexer) (Flow, error) {
	log.Debug().Msg("enter LexTemplate")
	defer log.Debug().Msg("exit LexTemplate")

	// terminal
	if !l.Available() {
		return nil, nil
	}

	// value := l.RunesAhead()

	// test if it's expression
	// then assignment
	// then comment
	// then text

	return nil, nil
}
