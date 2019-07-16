package value

var (
	emptyAbstractMap = make(AbstractMap)
)

type AbstractMap map[string]interface{}

func EmptyAbstractMap() AbstractMap {
	return emptyAbstractMap
}

func (m AbstractMap) IsMap() bool {
	return true
}

func (m AbstractMap) AsMap() AbstractMap {
	return m
}

func (m AbstractMap) UnsafeMap() AbstractMap {
	return m
}
