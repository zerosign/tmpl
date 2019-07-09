package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"log"
)

// LexText: lexing general text outside block template grammar.
//
//
func LexText(l base.Lexer) (Flow, error) {
	log.Println("enter logText")
	defer log.Println("leaving logText")

	for {
		value := l.RunesAhead()
		cursor := l.Cursor()

		if runes.HasPrefix(value, token.BlockExprOpen) {
			if cursor.IsValid() {
				l.Emit(token.TokenText)
			}
			return LexBlockExprOpen, nil
		} else if runes.HasPrefix(value, token.BlockCommentOpen) {
			if cursor.IsValid() {
				l.Emit(token.TokenBlockCommentOpen)
			}
			return LexBlockCommentOpen, nil
		} else if runes.HasPrefix(value, token.BlockAssignOpen) {
			if cursor.IsValid() {
				l.Emit(token.TokenBlockAssignOpen)
			}
			return LexBlockAssignOpen, nil
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	// emit current token type (Token)
	// since it's eof
	l.Emit(token.TokenText)

	// stop the lexer loop
	return nil, nil
}
