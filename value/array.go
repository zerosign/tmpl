package value

var (
	emptyAbstractArray = make(AbstractArray, 0)
)

type AbstractArray []interface{}

func EmptyAbstractArray() AbstractArray {
	return emptyAbstractArray
}
