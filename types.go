package tmpl

import (
	"fmt"
	"reflect"
	"runtime"
)

type StateFn func(*Lexer) (StateFn, error)

func (s StateFn) String() string {
	return fmt.Sprintf("<func %s>", runtime.FuncForPC(reflect.ValueOf(s).Pointer()).Name())
}
