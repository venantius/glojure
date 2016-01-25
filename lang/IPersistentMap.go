package lang

// TODO: extends Iterable, not sure how to handle in Go
type IPersistentMap interface {
	Associative

	AssocEx(interface{}, interface{}) IPersistentMap
	Without(interface{}) IPersistentMap
}
