package token

// Constant variable for each Token type.
//
// We don't model error type as TokenType since we have an explicit
// error value return for each Flow.
//
const (
	TokenNil Type = iota // since zero value of int is zero
	TokenText
	TokenBlockExprOpen
	TokenBlockExprClose
	TokenBlockAssignOpen
	TokenBlockAssignClose
	TokenBlockComment
	TokenBlockCommentOpen
	TokenBlockCommentClose
	TokenIf
	TokenFor
	TokenIn
	TokenBraceOpen
	TokenBraceClose
	TokenIdent
	TokenLetter
	TokenColon
	TokenComma
	TokenDeclType
	TokenDo
	TokenFuncName
	TokenParaOpen
	TokenParaClose
	TokenInteger
	TokenQuote
)
