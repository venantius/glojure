package lang

type IPersistentCollection interface {
	// IObj TODO (maybe AObj ?)
	Counted
	Seqable

	// TODO: replace with Cons
	Cons(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool
}
