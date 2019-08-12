package function

import (
	"fmt"
	"github.com/zerosign/tmpl/ast"
	"github.com/zerosign/tmpl/value"
	"reflect"
)


// FunctionDecl : function representation for tmpl external (golang) function.
//
// We limit argument types & return type to be value.Kind rather than using
// reflect.Kind directly. This enable us to limit the scope of the function itself.
//
// Supports for custom type (Struct) are still being argued whether is needed or not.
//
type FunctionDecl struct {
	name   []rune
	params []value.Kind
	contextFlag bool
	varargs value.Kind
	inner  reflect.Value
	result value.Kind
}

func (f FunctionDecl) Call(context Context, params []value.Value) (value.Value, error) {
	// TODO: @zerosign
	return nil, nil
}

func (f FunctionDecl) Mutate(context Context, node ast.Node, params []value.Value) (value.Value, error) {
	// TODO: @zerosign
	return nil, nil
}

func TypeOf(ty reflect.Type) (value.Kind, error) {
	switch ty.Kind() {
	case reflect.Map:
		return value.Map, nil
	case reflect.Array:
		return value.Array, nil
	case reflect.Bool:
		return value.Bool, nil
	case reflect.Int | reflect.Int8 | reflect.Int16 | reflect.Int32 | reflect.Int64:
		return value.Int, nil
	case reflect.Float32 | reflect.Float64:
		return value.Double, nil
	default:
		return value.Error, UnsupportedType(ty.Kind())
	}
}

// ReturnType: fetch return type of the function call
//
func ReturnType(name string, fn reflect.Value, ty reflect.Type) (value.Kind, error) {

	if ty.NumOut() != 2 {
		return value.Error, UnsupportedSizeReturnType(name)
	}

	lhs := ty.Out(0)
	rhs := ty.Out(1)

	if IsError(rhs) {
		return value.Error, MissingErrorReturnType(name)
	}

	kind, err := TypeOf(lhs)

	return kind, err
}

// Params : fetch all parameters that compatible with value.Kind.
//
// Params of functions :
// - first param ~ *Context (optional)
// - other params ...
// - last param should be var args of ... interface{} (optional)
//
func Params(fn reflect.Value, ty reflect.Type) ([]value.Kind, bool, value.Kind, []error) {
	var params []value.Kind = make([]value.Kind, 0)
	var errors []error = make([]error, 0)
	var isContext bool = false
	var variadic value.Kind = value.Error
	var it int = 0

	size := ty.NumIn()

	if ty.NumIn() == 0 {
		return []value.Kind{}, isContext, variadic, []error{ZeroFunctionParam()}
	}

	// check whether first arg is context or not
	if IsContext(ty.In(0)) {
		isContext = true
		if ty.NumIn() == 1 {
			return []value.Kind{}, isContext, variadic, []error{MissingFunctionParams("context")}
		} else {
			// since first params is context
			it = 1
		}
	}

	// check whether a function contains variadic args or not
	if ty.IsVariadic() {
		if ty.NumIn() == 1 {
			return []value.Kind{}, isContext, variadic, []error{MissingFunctionParams("options")}
		} else {
			var err error


			// since it's variadic, see https://golang.org/src/reflect/type.go?s=1321:7470#L28
			variadic, err = TypeOf(ty.In(ty.NumIn() - 1).Elem())

			if err != nil {
				return []value.Kind{}, isContext, variadic, []error{err}
			}

			size -= 1
		}
	}

	for ii := it; ii < size; ii++ {
		param := ty.In(ii)

		kind, err := TypeOf(param)

		if err != nil {
			errors = append(errors, err)
		}

		params = append(params, kind)
	}

	return params, isContext, variadic, errors
}

// Convert : function to convert given golang function to tmpl FunctionDecl
//
//
func Convert(name string, fptr interface{}) (*FunctionDecl, []error) {
	fn := reflect.ValueOf(fptr)
	ty := fn.Type()

	if ty.Kind() != reflect.Func {
		return nil, []error{TypeNotFunction(fn.Type().Name())}
	}

	// fetch all params from the function
	params, isContext, variadic, errors := Params(fn, ty)

	if len(errors) != 0 {
		return nil, errors
	}

	returnType, err := ReturnType(name, fn, ty)

	if err != nil {
		return nil, []error{err}
	}

	return &FunctionDecl{
		name: []rune(name),
		params: params,
		contextFlag: isContext,
		varargs: variadic,
		inner: fn,
		result: returnType,
	}, nil
}

func (f FunctionDecl) String() string {
	return fmt.Sprintf("<function:%s>", string(f.name))
}
