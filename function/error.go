package function

import (
	"fmt"
)

func TypeNotFunction(name string) error {
	fmt.Errorf("error type %s not a function", name)
}
