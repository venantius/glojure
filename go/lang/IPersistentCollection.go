package lang

type IPersistentCollection interface {
	Counted
	Seqable

	// TODO: replace with Cons
	Cons(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool
}
