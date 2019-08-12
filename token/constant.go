package token

// Constant variable for each Token type.
//
// We don't model error type as TokenType since we have an explicit
// error value return for each Flow.
//
const (
	Nil Type = iota // since zero value of int is zero
	Text
	If
	For
	In

	Ident
	Letter
	Colon
	Comma
	DeclType
	Do
	FuncName
	Integer
	Quote
	Escape
	String
	Double
	Sign
	Digit

	OpenBrace
	CloseBrace

	OpenPara
	ClosePara

	OpenBracket
	CloseBracket

	OpenExpr
	CloseExpr

	OpenAssign
	CloseAssign

	OpenComment
	CloseComment

	Comment

)
