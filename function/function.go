package function

import (
	"fmt"
	"github.com/zerosign/tmpl/value"
	"reflect"
)

type FunctionDecl struct {
	name   []rune
	params []value.Kind
	inner  reflect.Value
	result value.Kind
}

func (f FunctionDecl) Call(context Context, params []value.Value) (value.Value, error) {
	return nil, nil
}

func Params(fn reflect.Value, ty reflect.Type) []value.Kind {
	var params []value.Kind = make([]value.Kind, 0)

	size := ty.NumIn()

	for ii := 0; ii < size; ii++ {
		param := ty.In(ii)
		params = append(params, param)
	}

	return params
}

func Convert(fptr interface{}) (*FunctionDecl, error) {
	fn := reflect.ValueOf(fptr)
	ty := fn.Type()

	if ty.Kind() != reflect.Func {
		return nil, TypeNotFunction(fn.Type().Name())
	}

	// params := Params(fn, ty)

	// if ty.NumOut() != 2 {
	// 	// return
	// }

	return nil, nil
}

func (f *FunctionDecl) String() string {
	return fmt.Sprintf("<function:%s(%s)>")
}
