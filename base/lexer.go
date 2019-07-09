package base

import (
	"github.com/zerosign/tmpl/token"
	"io"
)

type Combinators interface {
	TakeWhile(conds ...token.RuneCond)
	AcceptAll(conds ...token.RuneCond) bool
	AcceptUntil(conds ...token.RuneCond) bool
	Accept(cond token.RuneCond) bool
	Ignore(cond token.RuneCond)
}

type LookAheadPtr interface {
	HasNext() bool
	Next() (token.Token, error)
	Advance()
	Available() bool
	RunesAhead() []rune
	CurrentRune() rune
}

type Emitter interface {
	Emit(t token.Type)
}

type HasCursor interface {
	Cursor() Cursor
	CursorMut() *Cursor
}

type Lexer interface {
	HasCursor
	Emitter
	LookAheadPtr
	Combinators
	io.Closer
}
