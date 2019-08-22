package lexer

import (
	"fmt"
	"github.com/zerosign/tmpl/common"
	"github.com/zerosign/tmpl/token"
	"io"
	"reflect"
	"runtime"
)

// Combinators : interface that deals with lexer combinators
//
type Combinators interface {
	TakeWhile(conds ...RuneCond)
	AcceptAll(conds ...RuneCond) bool
	Accept(cond RuneCond) bool
	Ignore(cond RuneCond)
}

// RuneCursorInfo : interface that deals with rune user info based on cursor info.
//
type RuneCursorInfo interface {
	CurrentRune() rune
	LastRune() (rune, error)
	StartRune() rune
}

// LookAheadPtr : interface that deals with look ahead operations
//
type LookAheadPtr interface {
	HasNext() bool
	Next() (token.Token, error)
	Advance()
	Available() bool
	RunesAhead() []rune
}

// TokenEmiter : interface that deals with emitting token
//
type TokenEmiter interface {
	Emit(tt token.Type)
}

// HasCursor : interface that represents mutator and accessor for cursor.
//
type HasCursor interface {
	Cursor() common.Cursor
	CursorMut() *common.Cursor
}

// Lexer : final interface that represents a lexer.
//
type Lexer interface {
	HasCursor
	TokenEmiter
	LookAheadPtr
	Combinators
	RuneCursorInfo
	io.Closer
	IsClosed() bool
}

// Flow: state function that represents control flow graph for the lexer.
//
type Flow func(Lexer) (Flow, error)

// String: string representation of flow function.
//
// It uses runtime & reflect from golang (it will be quite slow)
//
func (f Flow) String() string {
	return fmt.Sprintf("<func %s>", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
}
