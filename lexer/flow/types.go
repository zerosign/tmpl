package flow

import (
	"fmt"
	"github.com/zerosign/tmpl/base"
	"reflect"
	"runtime"
)

// Flow: state function that represents control flow graph for the lexer.
//
type Flow func(base.Lexer) (Flow, error)

// String: string representation of flow function.
//
// It uses runtime & reflect from golang (it will be quite slow)
//
func (f Flow) String() string {
	return fmt.Sprintf("<func %s>", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
}
