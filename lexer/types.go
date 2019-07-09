package lexer

import (
	"fmt"
	"reflect"
	"runtime"
)

// StateFn: state function that represents control flow graph for the lexer.
//
type StateFn func(*Lexer) (StateFn, error)

// String: string representation of state function.
//
// It uses runtime & reflect from golang (it will be quite slow)
//
func (s StateFn) String() string {
	return fmt.Sprintf("<func %s>", runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name())
}
