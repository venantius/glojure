package lang

type IPersistentCollection interface {
	Counted
	Seqable

	Cons(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool
}
