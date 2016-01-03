package lang

type IPersistentCollection interface {
	Counted
	Seqable

	// TODO: replace with Cons
	ConsIPersistentCollection(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool
}
