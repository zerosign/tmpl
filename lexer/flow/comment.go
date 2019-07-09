package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"log"
)

// LexBlockCommentOpen: lexing start of the comment block
//
//
func LexBlockCommentOpen(l base.Lexer) (Flow, error) {
	log.Println("enter logBlockCommentOpen")
	l.CursorMut().Incr(len(token.BlockCommentOpen))
	l.Emit(token.TokenBlockCommentOpen)
	return LexCommentBlock, nil
}

// LexCommentBlock: lexing the entire comment block
//
// block region contains :
// - consume everything until found BlockCommentClose
//
func LexCommentBlock(l base.Lexer) (Flow, error) {
	log.Println("enter logCommentBlock")
	for {
		value := l.RunesAhead()
		if runes.HasPrefix(value, token.BlockCommentClose) {
			l.Emit(token.TokenBlockComment)
			return LexCommentBlockClose, nil
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}
	return nil, NoMatchToken("comment", [][]rune{token.BlockCommentClose})
}

// LexCommentBlockClose: lexing close block of the comment block
//
//
func LexCommentBlockClose(l base.Lexer) (Flow, error) {
	log.Println("enter logCommentBlockClose")
	l.CursorMut().Incr(len(token.BlockCommentClose))
	l.Emit(token.TokenBlockCommentClose)
	return LexText, nil
}
