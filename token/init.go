package token

// Type : newtype enum to represents Token type.
//
type Type int

// String : string representation of token.Type.
//
//
func (tt Type) String() string {
	return Of(tt)
}

var (
	tokens = map[Type]string{
		Nil:          "<nil>",
		Text:         "<text>",
		OpenExpr:     "<block_expr_open>",
		CloseExpr:    "<block_expr_close>",
		OpenAssign:   "<block_assign_open>",
		CloseAssign:  "<block_assign_close>",
		Comment:      "<block_comment>",
		OpenComment:  "<block_comment_open>",
		CloseComment: "<block_comment_close>",
		If:           "<keyword_if>",
		For:          "<keyword_for>",
		In:           "<keyword_in>",
		Do:           "<do>",
		OpenBrace:    "<block_brace_open>",
		CloseBrace:   "<block_brace_close>",
		Ident:        "<ident>",
		Letter:       "<letter>",
		DeclType:     "<decl_type>",
	}
)

// Of : get string representation of token.
//
//
func Of(ty Type) string {
	val := tokens[ty]
	return val
}
