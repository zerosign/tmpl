package flow

import (
	"github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"unicode"
)

// LexFunctionCall
//
// call()
// call("")
// call(1)
// call(data)
// call(data, 1, "")
//
func LexFunctionCall(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter FunctionCall")
	defer log.Debug().Msg("exit FunctionCall")

	var err error

	// take while there is letter & lower
	l.TakeWhile(unicode.IsLetter, unicode.IsLower, token.IsSymbol)
	l.Emit(token.FuncName)

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	if l.CurrentRune() == runes.OpenPara {
		l.Emit(token.OpenPara)
	}

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	_, err = LexFunctionArgs(l)

	if err != nil {
		return nil, err
	}

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	if l.CurrentRune() == runes.ClosePara {
		l.Emit(token.ClosePara)
	}

	return nil, nil
}

//
// "data"
// 1
// data
// 0.0
// Value | (1 + (1 + 1) + 1)
//
func LexFunctionArgs(l base.Lexer) (Flow, error) {
	log.Debug().Msg("enter FunctionArgs")
	defer log.Debug().Msg("exit FunctionArgs")

	for {
		l.Ignore(token.IsWhitespace)
		// TODO(@zerosign): expression lexer need to be done first

		// ignore expression lexer first

	}

	return nil, nil
}

//
// Only can be used for function that returns the block or
// basically we could say it's function that have a block content
// inside of it (as template).
//
func LexFunctionDecl(l base.Lexer) (Flow, error) {
	return nil, nil
}
