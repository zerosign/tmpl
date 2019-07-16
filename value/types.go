package value

type Kind int

const (
	Null Kind = iota
	Bool
	Int
	Double
	Str
	Array
	Map
)

type Primitive interface {
	IsPrimitive() bool
}

type AsBool interface {
	IsBool() bool
	AsBool() (bool, error)
	UnsafeBool() bool
}

type AsInt interface {
	IsInt() bool
	AsInt() (int, error)
	UnsafeInt() int
}

type AsDouble interface {
	IsDouble() bool
	AsDouble() (float64, error)
	UnsafeDouble() float64
}

type AsStr interface {
	IsStr() bool
	AsStr() (string, error)
	UnsafeStr() string
}

type AsMap interface {
	IsMap() bool
	AsMap() (AbstractMap, error)
	UnsafeMap() AbstractMap
}

type AsArray interface {
	IsArray() bool
	AsArray() (AbstractArray, error)
	UnsafeArray() AbstractArray
}

type Value interface {
	Primitive
	AsBool
	AsInt
	AsDouble
	AsStr
	AsArray
	AsMap
}
