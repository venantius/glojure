package lang

type Associative interface {
	ILookup
	IPersistentCollection

	ContainsKey(key interface{}) bool
	EntryAt(key interface{}) IMapEntry
	Assoc(key interface{}, val interface{}) Associative
}
