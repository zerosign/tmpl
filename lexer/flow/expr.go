package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes/block"
	"github.com/zerosign/tmpl/runes/keyword"
	rutil "github.com/zerosign/tmpl/runes/util"
	"github.com/zerosign/tmpl/token"
)

func LexBlockExprOpen(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter LexBlockExprOpen")
	defer log.Debug().Msg("exit LexBlockExprOpen")

	l.CursorMut().Incr(len(block.OpenExpr))
	l.Emit(token.OpenExpr)
	return LexBlockExpr, nil
}

//
// block region contains :
// - for-statement
// - if-statement
//
func LexBlockExpr(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter LexBlockExpr")
	defer log.Debug().Msg("exit LexBlockExpr")

	for {
		// skip whitespace
		l.Ignore(token.IsWhitespace)

		value := l.RunesAhead()

		if rutil.HasPrefix(value, keyword.For) {
			// found keyword for
			l.Emit(token.For)

			return nil, nil
			// TODO: zerosign, return LexForStatement, nil
		} else if rutil.HasPrefix(value, keyword.If) {
			// found keyword if
			l.Emit(token.If)
			return LexStmtIf, nil
		}
	}

	// return nil, NoMatchToken("block", [][]rune{token.KeywordFor, token.KeywordIf})
}
