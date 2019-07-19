package function

import (
	"github.com/zerosign/tmpl/ast"
	"github.com/zerosign/tmpl/value"
)

// Function : interface that represents function
//
type Function interface {
	Name() []rune
	IsVoid() bool
	String() string
}

// Mutator : a mutator can fecth current ast.Node
//
// by providing ast.Node, one could query its parent
// and children ast.Node. This give the function a way
// to know its context. (things like yaml or json expander could use this
// so that it can be expand correctly, yaml need tab or space size).
//
type Mutator interface {
	Mutate(context Context, node ast.Node, params []value.Value) (value.Value, error)
}

// Void : a function that doesn't return a value to callee, but rather
//        only return error if fails
//
type Void interface {
	Exec(context Context, params []value.Value) error
}

// Call : a function that returns a value and error
//
type Call interface {
	Call(context Context, params []value.Value) (value.Value, error)
}
