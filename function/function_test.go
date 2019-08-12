package function

import (
	"fmt"
	"testing"
)

func Sample(ctx *Context, name string, options ...int) (string, error) {

	return fmt.Sprintf("hello world"), nil
}

func TestDeclareSimpleFunction(t *testing.T) {

}
