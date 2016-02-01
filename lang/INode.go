package lang

type INode interface {
	Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode
	Without(shift int, hash int, key interface{}) INode

	// Also returns interface{}
	Find(shift int, hash int, key interface{}, notFound interface{}) interface{}
	NodeSeq() ISeq

	AssocWithEdit(
		// bool for edit instead of AtomicReference, as with other transients
		edit bool,
		shift int,
		hash int,
		key interface{},
		val interface{},
		addedLeaf *Box,
	) INode

	WithoutWithEdit(
		edit bool,
		shift int,
		hash int,
		key interface{},
		removedLeaf *Box,
	) INode

	KVReduce(f IFn, init interface{}) interface{}
	Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{}

	Iterator(f IFn) *Iterator // I still do'nt know about this
}
