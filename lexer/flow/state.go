package flow

// //
// // for { key : String , value : Value } in (expr) do
// //
// func lexForStatement(l *lexer.Lexer) (lexer.Flow, error) {

// 	for {
// 		l.Ignore(IsWhitespace)

// 		if l.Accept(IsBraceOpen) {
// 			l.Emit(TokenBraceOpen)
// 			return lexBraceOpen, nil
// 		}
// 	}

// 	return nil, NoMatchToken("for-statement", [][]rune{[]rune{BraceOpen}})
// }

// func lexBraceOpen(l *lexer.Lexer) (lexer.Flow, error) {

// 	for {
// 		l.Ignore(IsWhitespace)
// 		value := l.RunesAhead()

// 		log.Printf("<ident>: %s\n", string(value))
// 	}

// 	return nil, nil
// }

// //
// // Lexing declaration
// //
// // <declaration> ::= <ident> (<token_colon> <type_decl>)? (<token_comma> <declaration>)?
// //
// func lexDeclaration(l *lexer.Lexer) (lexer.Flow, error) {

// 	lexIdent(l)

// 	if lexer.HasPrefix(l.RunesAhead(), []rune{SymbolColon}) {
// 		l.Accept(IsSymbolColon)
// 		// TODO: zerosign
// 	}

// 	return nil, nil
// }

// //
// // <type_decl> ::= <uppercase_letter> (<lowercase_letter>)*
// //
// func lexTypeDecl(l *lexer.Lexer) (lexer.Flow, error) {

// 	// <uppercase_letter>
// 	if !l.AcceptAll(unicode.IsLetter, unicode.IsUpper) {
// 		return nil, NotA("<uppercase_letter>")
// 	}

// 	// <lowercase_letter>*
// 	l.TakeWhile(unicode.IsLetter, unicode.IsLower)
// 	// l.AcceptUntil(unicode.IsLetter, unicode.IsLower)

// 	// skip whitespace
// 	l.Ignore(IsWhitespace)
// 	return nil, nil
// }

// //
// // <ident>         ::= <letter> (<integer> | <letter> | <symbol>)*
// // next token whitespace ':'
// //
// func lexIdent(l *lexer.Lexer) (lexer.Flow, error) {

// 	// only accept letter at beginning
// 	if !l.Accept(unicode.IsLetter) {
// 		return nil, NotA("<letter>")
// 	}

// 	// take while (<integer> | <letter> | <symbol>)*
// 	l.TakeWhile(unicode.IsLetter, unicode.IsDigit, IsSymbol)
// 	l.Emit(TokenIdent)

// 	// skip whitespace
// 	l.Ignore(IsWhitespace)

// 	return nil, NotA("<letter>")
// }

// func lexIfStatement(l *lexer.Lexer) (lexer.Flow, error) {

// 	return nil, nil
// }
