package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

// LexBlockCommentOpen: lexing start of the comment block
//
//
func LexBlockCommentOpen(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter BlockCommentOpen")
	defer log.Debug().Msg("exit BlockCommentOpen")

	l.CursorMut().Incr(len(token.BlockCommentOpen))
	l.Emit(token.TokenBlockCommentOpen)
	return LexBlockComment, nil
}

// LexCommentBlock: lexing the entire comment block
//
// block region contains :
// - consume everything until found BlockCommentClose
//
func LexBlockComment(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter BlockComment")
	defer log.Debug().Msg("exit BlockComment")

	for {
		value := l.RunesAhead()

		if runes.HasPrefix(value, token.BlockCommentClose) {
			break
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	l.Emit(token.TokenBlockComment)

	return LexCommentBlockClose, nil
}

// LexCommentBlockClose: lexing close block of the comment block
//
//
func LexCommentBlockClose(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter CommentBlockClose")
	defer log.Debug().Msg("exit CommentBlockClose")

	l.CursorMut().Incr(len(token.BlockCommentClose))
	l.Emit(token.TokenBlockCommentClose)

	// check whether there is available cursor

	if l.Available() {
		return LexText, nil
	} else {
		return nil, nil
	}
}
