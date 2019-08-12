package token

import (
	"github.com/zerosign/tmpl/runes"
	"unicode"
)

type _void struct{}

var (
	void = _void{}

	whitespace = map[rune]_void{
		' ':  void,
		'\t': void,
		'\n': void,
		'\r': void,
	}

	symbol = map[rune]_void{
		'_':  void,
		'-':  void,
		'\'': void,
	}

	// RuneCond functions for several runes
	IsOpenBrace    = IsRuneEq(runes.OpenBrace)
	IsCloseBrace   = IsRuneEq(runes.CloseBrace)
	IsOpenPara     = IsRuneEq(runes.OpenPara)
	IsClosePara    = IsRuneEq(runes.ClosePara)
	IsOpenBracket  = IsRuneEq(runes.OpenBracket)
	IsCloseBracket = IsRuneEq(runes.CloseBracket)
	IsColon        = IsRuneEq(runes.Colon)
	IsQuote        = IsRuneEq(runes.Quote)
	IsDot          = IsRuneEq(runes.Dot)

	IsPrimitive = func(ch rune) bool {
		return IsQuote(ch) && unicode.IsDigit(ch)
	}
)
