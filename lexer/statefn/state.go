package statefn

import (
	"github.com/zerosign/tmpl/lexer"
	"log"
	"unicode"
)

func lexBlockExprOpen(l *Lexer) (StateFn, error) {
	log.Println("enter logBlockExprOpen")

	l.current += len(BlockExprOpen)
	l.Emit(TokenBlockExprOpen)
	return lexBlock, nil
}

func lexBlockAssignOpen(l *Lexer) (StateFn, error) {
	log.Println("enter logBlockAssignOpen")

	l.current += len(BlockAssignOpen)
	l.Emit(TokenBlockAssignOpen)
	return lexAssignBlock, nil
}

//
// block region contains :
// - for-statement
// - if-statement
// - template-decl-statement
//
func lexBlock(l *Lexer) (StateFn, error) {
	log.Println("enter logBlock")
	for {
		// skip whitespace
		l.Ignore(IsWhitespace)

		value := l.RunesAhead()

		if HasPrefix(value, KeywordFor) {
			// found keyword for
			l.Emit(TokenFor)
			return lexForStatement, nil
		} else if HasPrefix(value, KeywordIf) {
			// found keyword if
			l.Emit(TokenIf)
			return lexIfStatement, nil
		}
	}

	return nil, NoMatchToken("block", [][]rune{KeywordFor, KeywordIf})
}

//
// for { key : String , value : Value } in (expr) do
//
func lexForStatement(l *Lexer) (StateFn, error) {

	for {
		l.Ignore(IsWhitespace)

		if l.Accept(IsBraceOpen) {
			l.Emit(TokenBraceOpen)
			return lexBraceOpen, nil
		}
	}

	return nil, NoMatchToken("for-statement", [][]rune{[]rune{BraceOpen}})
}

func lexBraceOpen(l *Lexer) (StateFn, error) {

	for {
		l.Ignore(IsWhitespace)
		value := l.RunesAhead()

		log.Printf("<ident>: %s\n", string(value))
	}

	return nil, nil
}

//
// Lexing declaration
//
// <declaration> ::= <ident> (<token_colon> <type_decl>)? (<token_comma> <declaration>)?
//
func lexDeclaration(l *Lexer) (StateFn, error) {

	lexIdent(l)

	if HasPrefix(l.RunesAhead(), []rune{SymbolColon}) {
		l.Accept(IsSymbolColon)
		// TODO: zerosign
	}

	return nil, nil
}

//
// <type_decl> ::= <uppercase_letter> (<lowercase_letter>)*
//
func lexTypeDecl(l *Lexer) (StateFn, error) {

	// <uppercase_letter>
	if !l.AcceptAll(unicode.IsLetter, unicode.IsUpper) {
		return nil, NotA("<uppercase_letter>")
	}

	// <lowercase_letter>*
	l.TakeWhile(unicode.IsLetter, unicode.IsLower)
	// l.AcceptUntil(unicode.IsLetter, unicode.IsLower)

	// skip whitespace
	l.Ignore(IsWhitespace)
	return nil, nil
}

//
// <ident>         ::= <letter> (<integer> | <letter> | <symbol>)*
// next token whitespace ':'
//
func lexIdent(l *Lexer) (StateFn, error) {

	// only accept letter at beginning
	if !l.Accept(unicode.IsLetter) {
		return nil, NotA("<letter>")
	}

	// take while (<integer> | <letter> | <symbol>)*
	l.TakeWhile(unicode.IsLetter, unicode.IsDigit, IsSymbol)
	l.Emit(TokenIdent)

	// skip whitespace
	l.Ignore(IsWhitespace)

	return nil, NotA("<letter>")
}

func lexIfStatement(l *Lexer) (StateFn, error) {

	return nil, nil
}

//
// block region contains :
// - expression statement
//
func lexAssignBlock(l *Lexer) (StateFn, error) {
	log.Println("enter logAssignBlock")
	return nil, nil
}
