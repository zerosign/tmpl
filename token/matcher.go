package token

import (
	"unicode"
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

	Quote rune = '"'

	Colon rune = ':'

	Comma rune = ','
)

var (
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
	KeywordDo    = []rune("do")
	KeywordEnd   = []rune("end")
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
	IsQuote        = IsRuneEq(Quote)

	IsPrimitive = func(ch rune) bool {
		return IsQuote(ch) && unicode.IsDigit(ch)
	}
)
