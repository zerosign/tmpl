package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

func LexBlockExprOpen(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter LexBlockExprOpen")
	defer log.Debug().Msg("exit LexBlockExprOpen")

	l.CursorMut().Incr(len(token.BlockExprOpen))
	l.Emit(token.TokenBlockExprOpen)
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

		if runes.HasPrefix(value, token.KeywordFor) {
			// found keyword for
			l.Emit(token.TokenFor)

			return nil, nil
			// TODO: zerosign, return LexForStatement, nil
		} else if runes.HasPrefix(value, token.KeywordIf) {
			// found keyword if
			l.Emit(token.TokenIf)
			return LexStmtIf, nil
		}
	}

	return nil, NoMatchToken("block", [][]rune{token.KeywordFor, token.KeywordIf})
}
