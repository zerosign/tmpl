package function

import (
	"fmt"
	"testing"
)

func Sample(ctx Context) (string, error) {
	return fmt.Sprintf("hello world"), nil
}

func TestDeclareSimpleFunction() {
	fn := Convert(Sample)
}
