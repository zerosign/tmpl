package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes/keyword"
	rutil "github.com/zerosign/tmpl/runes/util"
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

	l.CursorMut().Incr(len(keyword.For))
	l.Emit(token.For)

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

	if rutil.HasPrefix(l.RunesAhead(), keyword.Do) {

		l.CursorMut().Incr(len(keyword.Do))
		l.Emit(token.Do)

	} else {
		return nil, NotA(token.Do)
	}

	return nil, nil
}

func LexStmtForIn(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter StmtForIn")
	defer log.Debug().Msg("exit StmtForIn")

	l.CursorMut().Next()
	l.Emit(token.If)

	return nil, nil
}
