package statefn

import (
	"github.com/zerosign/tmpl/lexer"
	"log"
)

// LexBlockCommentOpen: lexing start of the comment block
//
//
func LexBlockCommentOpen(l *lexer.Lexer) (StateFn, error) {
	log.Println("enter logBlockCommentOpen")
	l.current += len(token.BlockCommentOpen)
	l.Emit(TokenBlockCommentOpen)
	return LexCommentBlock, nil
}

// LexCommentBlock: lexing the entire comment block
//
// block region contains :
// - consume everything until found BlockCommentClose
//
func LexCommentBlock(l *Lexer) (StateFn, error) {
	log.Println("enter logCommentBlock")
	for {
		value := l.inner[l.current:]
		if HasPrefix(value, token.BlockCommentClose) {
			l.Emit(token.TokenBlockComment)
			return LexCommentBlockClose, nil
		}

		if l.Available() {
			l.Advance()
		} else {
			break
		}
	}
	return nil, NoMatchToken("comment", [][]rune{BlockCommentClose})
}

// LexCommentBlockClose: lexing close block of the comment block
//
//
func LexCommentBlockClose(l *Lexer) (StateFn, error) {
	log.Println("enter logCommentBlockClose")
	l.current += len(BlockCommentClose)
	l.Emit(TokenBlockCommentClose)
	return LexText, nil
}
