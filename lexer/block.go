package lexer

var (
	// blocks
	OpenStmt     = []rune("{{")
	CloseStmt    = []rune("}}")
	OpenExpr     = []rune("{=")
	CloseExpr    = []rune("=}")
	OpenComment  = []rune("{#")
	CloseComment = []rune("#}")
)
