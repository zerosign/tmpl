package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

func LexIfStatement(l base.Lexer) (Flow, error) {

	for {
		l.Ignore(token.IsWhitespace)

		lexeme := l.RunesAhead()

		if runes.HasPrefix(lexeme, token.KeywordIf) {

		}
	}

	return nil, nil
}
