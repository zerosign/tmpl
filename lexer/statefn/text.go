package statefn

import (
	"github.com/zerosign/tmpl/lexer"
	"github.com/zerosign/tmpl/lexer/token"
	"log"
	"unicode"
)

// LexText: lexing general text outside block template grammar.
//
//
func LexText(l *lexer.Lexer) (lexer.StateFn, error) {
	log.Println("enter logText")
	defer log.Println("leaving logText")

	for {
		value := l.RunesAhead()

		if HasPrefix(value, token.BlockExprOpen) {
			if l.current > l.start {
				l.Emit(TokenText)
			}
			return LexBlockExprOpen, nil
		} else if HasPrefix(value, token.BlockCommentOpen) {
			if l.current > l.start {
				l.Emit(token.TokenBlockCommentOpen)
			}
			return LexBlockCommentOpen, nil
		} else if HasPrefix(value, token.BlockAssignOpen) {
			if l.current > l.start {
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
	l.Emit(TokenText)

	// stop the lexer loop
	return nil, nil
}
