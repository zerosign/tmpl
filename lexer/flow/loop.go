package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
)

//
// for { key: String, value: Value } in rlookup("vault", "") do
//
//
func LexStmtFor(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter ForStatement")
	defer log.Debug().Msg("exit ForStatement")

	var err error

	l.CursorMut().Incr(len(token.KeywordFor))
	l.Emit(token.TokenFor)

	_, err = LexBlockDecl(l)

	if err != nil {
		return nil, err
	}

	l.Ignore(token.IsWhitespace)

	_, err = LexStmtForIn(l)

	if err != nil {
		return nil, err
	}

	l.Ignore(token.IsWhitespace)

	_, err = LexFunctionCall(l)

	if err != nil {
		return nil, err
	}

	l.Ignore(token.IsWhitespace)

	if runes.HasPrefix(l.RunesAhead(), token.KeywordDo) {

		l.CursorMut().Incr(len(token.KeywordDo))
		l.Emit(token.TokenDo)

	} else {
		return nil, NotA(token.TokenDo)
	}

	return nil, nil
}

func LexStmtForIn(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter StmtForIn")
	defer log.Debug().Msg("exit StmtForIn")

	l.CursorMut().Next()
	l.Emit(token.TokenIf)

	return nil, nil
}
