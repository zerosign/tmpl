package grammar

import (
	"fmt"
)

type RuneCond func(ch rune) bool
type TokenType int

const (
	TokenError TokenType = iota
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
	TokenBraceOpen
	TokenIdent
)

var (
	tokens = map[TokenType]string{
		TokenError:             "<error>",
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

	whitespace = map[rune]struct{}{
		' ':  struct{}{},
		'\t': struct{}{},
		'\n': struct{}{},
		'\r': struct{}{},
	}

	symbol = map[rune]struct{}{
		'_':  struct{}{},
		'-':  struct{}{},
		'\'': struct{}{},
	}

	// blocks
	BlockExprOpen     = []rune("{{")
	BlockExprClose    = []rune("}}")
	BlockAssignOpen   = []rune("{=")
	BlockAssignClose  = []rune("=}")
	BlockCommentOpen  = []rune("{#")
	BlockCommentClose = []rune("#}")

	// reserved keywords
	KeywordFor   = []rune("for")
	KeywordIf    = []rune("if")
	KeywordElse  = []rune("else")
	KeywordElsif = []rune("elsif")

	// symbol
	SymbolColon rune = ':'

	// RuneCond functions for several runes
	IsBraceOpen    = IsRuneEq(BraceOpen)
	IsBraceClose   = IsRuneEq(BraceClose)
	IsParaOpen     = IsRuneEq(ParaOpen)
	IsParaClose    = IsRuneEq(ParaClose)
	IsBracketOpen  = IsRuneEq(BracketOpen)
	IsBracketClose = IsRuneEq(BracketClose)
	IsSymbolColon  = IsRuneEq(SymbolColon)
)

const (

	// brace
	BraceOpen  rune = '{'
	BraceClose rune = '}'

	// para
	ParaOpen  rune = '('
	ParaClose rune = ')'

	// bracket
	BracketOpen  rune = '['
	BracketClose rune = ']'
)

func (tt TokenType) String() string {
	return fmt.Sprintf("%s", tokens[tt])
}

type Token struct {
	Type  TokenType
	Value []rune
}

func (t Token) String() string {
	return fmt.Sprintf("Token { %s, %s }", t.Type, string(t.Value))
}

func IsRuneEq(ch rune) RuneCond {
	return func(v rune) bool { return v == ch }
}

func IsWhitespace(ch rune) bool {
	_, flag := whitespace[ch]
	return flag
}

func IsSymbol(ch rune) bool {
	_, flag := symbol[ch]
	return flag
}
