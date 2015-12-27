package lang

type IPersistentCollection interface {
	Counted
	Seqable

	cons(o interface{}) IPersistentCollection
	empty() IPersistentCollection
	equiv(o interface{}) bool
}
