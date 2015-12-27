package lang

type Associative interface {
	ILookup
	IPersistentCollection

	containsKey(key interface{}) bool
	entryAt(key interface{}) IMapEntry
	assoc(key interface{}, val interface{}) Associative
}
