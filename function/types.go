package function

// Function : interface that represents function
//
//
type Function interface {
	Name() []rune
	IsVoid() bool
	String() string
}
