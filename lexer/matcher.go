package lexer

import (
	"unicode"
)

type _void struct{}

var (
	void            = _void{}
	newline         = map[rune]_void{'\r': void, '\n': void}
	whitespace_only = map[rune]_void{' ': void, '\t': void}
)

type RuneCond func(ch rune) bool

func runeCond(ch rune) RuneCond {
	return func(r rune) bool { return ch == r }
}

func and(lhs, rhs RuneCond) RuneCond {
	return func(r rune) bool { return lhs(r) && rhs(r) }
}

func or(lhs, rhs RuneCond) RuneCond {
	return func(r rune) bool { return lhs(r) || rhs(r) }
}

func isNot(cond RuneCond) RuneCond {
	return func(r rune) bool {
		return !cond(r)
	}
}

// IsNewline
//
//
func IsNewline(r rune) bool {
	_, ok := newline[r]
	return ok
}

// IsWhitespaceOnly
//
//
func IsWhitespaceOnly(r rune) bool {
	_, ok := whitespace_only[r]
	return ok
}

// IsWhitespace : translation of ebnf rule of `whitepace`
//
//
func IsWhitespace(r rune) bool {
	return IsNewline(r) || IsWhitespaceOnly(r)
}

// IsUppercase : translation of ebnf rule of 'uppercase_letter'
//
// it use unicode.IsUpper
//
func IsUppercase(r rune) bool {
	return unicode.IsUpper(r)
}

// IsLowercase : translation of ebnf rule of `lowercase_letter`
//
// it use unicode.IsLower
//
func IsLowercase(r rune) bool {
	return unicode.IsLower(r)
}

// IsDigit: translation of ebnf rule of 'digit'
//
// digit = "0" ... "9" .
//
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsDigitZero: translation of ebnf rule of `digit_zero`.
//
// digit_zero = "0" .
//
func IsDigitZero(r rune) bool {
	return r == '0'
}

// IsDigitWithoutZero: translation of ebnf rule of `digit_without_zero`
//
// digit_without_zero = "1"..."9" .
//
func IsDigitWithoutZero(r rune) bool {
	return !IsDigitZero(r) && IsDigit(r)
}

// IsAny: translation of ebnf rule of `any`
//
// any = "\u0000"..."\uffff" .
//
func IsAny(r rune) bool {
	return r >= '\u0000' && r <= '\uffff'
}
