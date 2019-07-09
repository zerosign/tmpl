package token

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
