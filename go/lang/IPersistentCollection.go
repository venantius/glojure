package lang

type IPersistentCollection interface {
	Counted
	Seqable

	ConsIPersistentCollection(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool
}
