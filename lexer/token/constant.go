package token

// Constant variable for each Token type.
//
// We don't model error type as TokenType since we have an explicit
// error value return for each StateFn.
//
const (
	TokenText Type = iota
	TokenBlockExprOpen
	TokenBlockExprClose
	TokenBlockAssignOpen
	TokenBlockAssignClose
	TokenBlockComment
	TokenBlockCommentOpen
	TokenBlockCommentClose
	TokenIf
	TokenFor
	TokenBraceOpen
	TokenIdent
)
