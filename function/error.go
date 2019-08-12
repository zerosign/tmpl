package function

import (
	"fmt"
	"reflect"
)

func ZeroFunctionParam() error {
	return fmt.Errorf("function paramater size shoudn't be 0")
}

func MissingFunctionParams(context string) error {
	return fmt.Errorf("function shouldn't only contains options or context as parameter: only %s", context)
}

func TypeNotFunction(name string) error {
	return fmt.Errorf("error type %s not a function", name)
}

func UnsupportedSizeReturnType(name string) error {
	return fmt.Errorf("unsupported function %s return type, return type should be at least 2", name)
}

func MissingErrorReturnType(name string) error {
	return fmt.Errorf("second (rhs) return type of the function %s should return an error", name)
}

func UnsupportedType(kind reflect.Kind) error {
	return fmt.Errorf("type %v are unsupported", kind)
}
