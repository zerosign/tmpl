package lexer

import (
	"github.com/rs/zerolog/log"
)

//
// Block = TextBlock | CommentBlock | StmtBlock | ExprBlock .
//
// template are also terminal state.
//
func BlockFlow(l Lexer) (Flow, error) {
	log.Debug().Msg("enter BlockFlow")
	defer log.Debug().Msg("exit BlockFlow")

	// terminal
	if !l.Available() {
		return nil, nil
	}

	value := l.RunesAhead()

	// check whether its comment block
	if IsCommentBlock(value) {
		return CommentBlockFlow, nil
	} else if IsStmtBlock(value) {
		return StmtBlockFlow, nil
	} else if IsExprBlock(value) {
		return ExprBlockFlow, nil
	} else {
		return TextBlockFlow, nil
	}
}
