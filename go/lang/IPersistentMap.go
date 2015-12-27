package lang

// TODO: extends Iterable, not sure how to handle in Go
type IPersistentMap interface {
	Associative

	assocEx(interface{}, interface{}) IPersistentMap
	without(interface{}) IPersistentMap
}
