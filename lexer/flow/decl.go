package flow

import (
	"github.com/zerosign/tmpl/base"
	// "github.com/zerosign/tmpl/runes"
	"github.com/zerosign/tmpl/token"
	"log"
	"unicode"
)

//
// { key: String, value: Value }
//
// This flow does terminal, any call to this need to
// be called & managed in call site.
//
func LexBlockDecl(l base.Lexer) (Flow, error) {
	log.Println("enter Declaration")
	defer log.Println("exit Declaration")

	value := l.CurrentRune()

	if value == token.BraceOpen {
		return LexBraceOpen, nil
	} else {
		return nil, NotA(token.TokenBraceOpen)
	}
}

//
// emit: `{`, next: LexDeclaration
//
func LexBraceOpen(l base.Lexer) (Flow, error) {
	log.Println("enter BraceOpen")
	defer log.Println("exit BraceOpen")

	l.CursorMut().Next()
	l.Emit(token.TokenBraceOpen)

	return LexDeclaration, nil
}

// LexDeclaration : lexing declaration
//
// ident : type (,)
//
// Valid string :
//
// - ident, ident, ..
// - ident : type, ident, ...
// - ident
// - ident : type
//
// it returns nil for the next state, since
// the next state are depends on the callsite
//
func LexDeclaration(l base.Lexer) (Flow, error) {
	log.Println("enter Declaration")
	defer log.Println("exit Declaration")

	var err error

	// fetch multiple declaration
	for {

		// ignore whitespace
		l.Ignore(token.IsWhitespace)

		// fetch ident
		_, err = LexIdent(l)

		if err != nil {
			return nil, err
		}

		// ignore whitespace
		l.Ignore(token.IsWhitespace)

		// check whether ident declare a type or not
		if l.CurrentRune() == token.Colon {

			l.Emit(token.TokenColon)

			// fetch type of ident
			_, err = LexDeclType(l)

			if err != nil {
				return nil, err
			}
		}

		// ignore whitespace
		l.Ignore(token.IsWhitespace)

		if l.CurrentRune() == token.Comma {
			// emit token Comma
			l.Emit(token.TokenComma)

			// since it has comma, then there should have next declaration
			l.Next()

		} else {
			// this break only when there were no 'token.Comma' found
			break
		}
	}

	return nil, nil
}

//
// LexTypeDecl : lexing type declaration
//
// a type should starts with uppercase letter
//
// returns nil, next state are being controlled in callsite
//
func LexDeclType(l base.Lexer) (Flow, error) {
	log.Println("enter TypeDecl")
	defer log.Println("exit TypeDecl")

	l.Ignore(token.IsWhitespace)

	// check whether current rune is uppercase utf8 character
	if !unicode.IsUpper(l.CurrentRune()) {
		return nil, CaseError(upperCase, lowerCase)
	}

	// check whether current rune is acceptable utf8 character
	// (letter | symbol)
	l.TakeWhile(unicode.IsLetter, token.IsSymbol)

	// emit current decl type
	l.Emit(token.TokenDeclType)

	return nil, nil
}

// emit: '}', next: callee site defined
//
func LexBraceClose(l base.Lexer) (Flow, error) {
	log.Println("enter BraceClose")
	defer log.Println("exit BraceClose")

	l.CursorMut().Next()
	l.Emit(token.TokenBraceClose)

	return nil, nil
}
