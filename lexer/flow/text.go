package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"log"
)

var (
	NextStateTokens = [][]rune{token.BlockExprOpen, token.BlockAssignOpen, token.BlockCommentOpen}
)

// LexText: lexing general text outside block template grammar.
//
//
func LexText(l base.Lexer) (Flow, error) {
	var value []rune

	log.Println("enter LexText")
	defer log.Println("exit LexText")

	for {
		value = l.RunesAhead()

		// scan for it siblings token
		if runes.HasAllPrefixes(value, NextStateTokens) {
			break
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	// emit last token that being scanned (in here token.TokenText)
	l.Emit(token.TokenText)

	return LexTemplate, nil
}
