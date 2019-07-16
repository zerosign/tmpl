package flow

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"log"
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
	log.Println("enter FunctionCall")
	defer log.Println("exit FunctionCall")

	var err error

	// take while there is letter & lower
	l.TakeWhile(unicode.IsLetter, unicode.IsLower, token.IsSymbol)
	l.Emit(token.TokenFuncName)

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	if l.CurrentRune() == token.ParaOpen {
		l.Emit(token.TokenParaOpen)
	}

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	_, err = LexFunctionArgs(l)

	if err != nil {
		return nil, err
	}

	// ignore whitespace
	l.Ignore(token.IsWhitespace)

	if l.CurrentRune() == token.ParaClose {
		l.Emit(token.TokenParaClose)
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
	log.Println("enter FunctionArgs")
	defer log.Println("exit FunctionArgs")

	for {
		l.Ignore(token.IsWhitespace)

		//

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
