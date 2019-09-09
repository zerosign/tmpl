package lexer

import (
	"fmt"
	"github.com/zerosign/tmpl/data"
	"github.com/zerosign/tmpl/token"
	"reflect"
	"runtime"
)

type HasCursor interface {
	CursorMut() *Cursor
	Cursor() Cursor
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
