package token

// IsRuneEq : Helper function that return RuneCond
//
// Example:
// var IsTab = IsRune('\t')
//
func IsRuneEq(ch rune) RuneCond {
	return func(v rune) bool { return v == ch }
}

// IsWhitespace : check whether a given rune is whitespace or not
//
//
func IsWhitespace(ch rune) bool {
	_, flag := whitespace[ch]
	return flag
}

// IsSymbol : check whether a given rune is allowed symbol or not
//
// see token.symbol for the complete lists of allowed symbols.
//
func IsSymbol(ch rune) bool {
	_, flag := symbol[ch]
	return flag
}
