package lang

// NOTE: Extends Associative, Sequential, IPersistentStack, Reversible, Indexed
type IPersistentVector interface {
	Sequential // Empty

	// ILookup
	ValAt(key interface{}, notFound interface{}) interface{}

	// IPersistentCollection
	Seqable

	ConsIPersistentCollection(o interface{}) IPersistentCollection
	Empty() IPersistentCollection
	Equiv(o interface{}) bool

	// IPersistentStack
	Peek() interface{}
	Pop() IPersistentStack

	Reversible
	Indexed

	ContainsKey(key interface{}) bool
	EntryAt(i interface{}) IMapEntry

	Length() int
	Assoc(key interface{}, val interface{}) Associative
	AssocN(i int, val interface{}) IPersistentVector
	ConsIPersistentVector(i interface{}) IPersistentVector
}
