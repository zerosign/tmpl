package token

// Type : newtype enum to represents Token type.
//
type Type int

// String : string representation of token.Type.
//
//
func (tt Type) String() string {
	return TokenOf(tt)
}

var (
	tokens = map[Type]string{
		TokenText:              "<text>",
		TokenBlockExprOpen:     "<block_expr_open>",
		TokenBlockExprClose:    "<block_expr_close>",
		TokenBlockAssignOpen:   "<block_assign_open>",
		TokenBlockAssignClose:  "<block_assign_close>",
		TokenBlockComment:      "<block_comment>",
		TokenBlockCommentOpen:  "<block_comment_open>",
		TokenBlockCommentClose: "<block_comment_close>",
		TokenIf:                "<keyword_if>",
		TokenFor:               "<keyword_for>",
		TokenBraceOpen:         "<block_brace_open>",
		TokenIdent:             "<ident>",
	}
)

// TokenOf : get string representation of token.
//
//
func TokenOf(ty Type) string {
	val := tokens[ty]
	return val
}
