package lang

type IMapIterable interface {
	KeyIterator() Iterator
	ValIterator() Iterator
}
