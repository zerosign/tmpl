package function

import (
	"reflect"
)

var (
	errorInterface = reflect.TypeOf((*error)(nil)).Elem()
	contextInterface = reflect.TypeOf((*Context)(nil)).Elem()
)

// IsError: check whether type is error or not
//
//
func IsError(other reflect.Type) bool {
	return other.Implements(errorInterface)
}

//
// IsContext: check whether type is context or not
//
// the given type should be direct instance of newtype Context
// rather than typeof `map[string]interface{}` directly.
//
func IsContext(other reflect.Type) bool {
	return other == contextInterface
}
