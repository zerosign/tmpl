package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes/block"
	rutil "github.com/zerosign/tmpl/runes/util"
	"github.com/zerosign/tmpl/token"
)

var (
	NextStateTokens = [][]rune{block.OpenExpr, block.OpenAssign, block.OpenComment}
)

// LexText: lexing general text outside block template grammar.
//
//
func LexText(l base.Lexer) (Flow, error) {
	var value []rune
	log.Debug().Msg("enter LexText")
	defer log.Debug().Msg("exit LexText")

	for {
		value = l.RunesAhead()

		// scan for it siblings token
		if rutil.AnyPrefixes(value, NextStateTokens) {
			break
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	// emit last token that being scanned (in here token.TokenText)
	l.Emit(token.Text)

	return LexTemplate, nil
}
