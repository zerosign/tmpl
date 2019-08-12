package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes/block"
	rutil "github.com/zerosign/tmpl/runes/util"
	"github.com/zerosign/tmpl/token"
)

// LexBlockCommentOpen: lexing start of the comment block
//
//
func LexBlockCommentOpen(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter BlockCommentOpen")
	defer log.Debug().Msg("exit BlockCommentOpen")

	l.CursorMut().Incr(len(block.OpenComment))
	l.Emit(token.OpenComment)
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

		if rutil.HasPrefix(value, block.CloseComment) {
			break
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	l.Emit(token.Comment)

	return LexCommentBlockClose, nil
}

// LexCommentBlockClose: lexing close block of the comment block
//
//
func LexCommentBlockClose(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter CommentBlockClose")
	defer log.Debug().Msg("exit CommentBlockClose")

	l.CursorMut().Incr(len(block.CloseComment))
	l.Emit(token.CloseComment)

	// check whether there is available cursor

	if l.Available() {
		return LexText, nil
	} else {
		return nil, nil
	}
}
