package function

import (
	"testing"
	"reflect"
	"fmt"
)

func TestIsError(t * testing.T) {
	var err = fmt.Errorf("Hello %v", "world")

	ty := reflect.TypeOf(err)

	if !IsError(ty) {
		t.Errorf("error '%v' should be expected as an error", err)
	}
}

func TestIsContext(t *testing.T) {
	var data = make(map[string]interface{}, 0)
	var context = Context{}
	var ty reflect.Type

	ty = reflect.TypeOf(data)

	if IsContext(ty) {
		t.Errorf("'%v' need to be a direct instance of newtype Context rather than its true alias form", data)
	}

	ty = reflect.TypeOf(context)

	if !IsContext(ty) {
		t.Errorf("'%v' should be direct instance of newtype Context", context)
	}
}
