package lexer

import (
	"github.com/zerosign/tmpl/token"
	"io"
)

// Combinators : interface that deals with lexer combinators
//
//
type Combinators interface {
	TakeWhile(conds ...token.RuneCond)
	AcceptAll(conds ...token.RuneCond) bool
	// AcceptUntil(conds ...token.RuneCond) bool
	Accept(cond token.RuneCond) bool
	Ignore(cond token.RuneCond)
}

// LookAheadPtr : interface that deals with look ahead pointer (cursor + runes)
//
//
type LookAheadPtr interface {
	HasNext() bool
	Next() (token.Token, error)
	Advance()
	Available() bool
	RunesAhead() []rune
	CurrentRune() rune
	PeekRune() rune
	LastRune() (rune, error)
	StartRune() rune
}

// Emitter : interface that deals with emitting token
//
//
type Emitter interface {
	Emit(t token.Type)
}

// HasCursor : interface that represents a mutator and getter for cursor.
//
//
type HasCursor interface {
	Cursor() Cursor
	CursorMut() *Cursor
}

// Lexer : final interface that represents a Lexer.
//
//
type Lexer interface {
	HasCursor
	Emitter
	LookAheadPtr
	Combinators
	io.Closer
	IsClosed() bool
}
