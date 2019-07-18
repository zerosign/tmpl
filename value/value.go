package value

import (
	"reflect"
)

func TypeValue(ty reflect.Type) Kind {

	switch ty.Kind() {
	case reflect.Int || reflect.Int8 || reflect.Int16 || reflect.Int32 || reflect.Int64:
		return Int
	default:
		return Unit
	}
	if ty.Kind() == reflect.Int {
		return Int
	}
}
