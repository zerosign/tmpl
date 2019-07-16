package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

func LexStmtIf(l base.Lexer) (Flow, error) {

	for {
		l.Ignore(token.IsWhitespace)

		lexeme := l.RunesAhead()

		if runes.HasPrefix(lexeme, token.KeywordIf) {

		}
	}

	return nil, nil
}

func LexStmtElsif(l base.Lexer) (Flow, error) {
	return nil, nil
}

func LexStmtElse(l base.Lexer) (Flow, error) {
	return nil, nil
}
