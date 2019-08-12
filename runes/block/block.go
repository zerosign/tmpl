package block

var (
	// blocks
	OpenExpr     = []rune("{{")
	CloseExpr    = []rune("}}")
	OpenAssign   = []rune("{=")
	CloseAssign  = []rune("=}")
	OpenComment  = []rune("{#")
	CloseComment = []rune("#}")
)
