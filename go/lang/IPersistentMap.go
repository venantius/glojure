package lang

// NOTE: Extends Iterable, Associative, Counted
type IPersistentMap interface {
	assoc(interface{}, interface{}) IPersistentMap
	assocEx(interface{}, interface{}) IPersistentMap
	without(interface{}) IPersistentMap
}
