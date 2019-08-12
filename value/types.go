package value

type Kind int

const (
	Error Kind = -1
	Unit  Kind = iota
	Bool
	Int
	Double
	Str
	Array
	Map
	Struct
)

type IsUnit interface {
	IsUnit() bool
}

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
	IsUnit
	Primitive
	AsBool
	AsInt
	AsDouble
	AsStr
	AsArray
	AsMap
}
