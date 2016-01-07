package lang

/*
	Note copied from JVM Clojure.

	A persistent rendition of Phil Bagwell's Hash Array Mapped Trie

	Uses path copying for persistence. HashCollision leaves vs. extended hashing
	Node polymorphism vs. conditionals
	No sub-tree pools or root-resizing
	Any errors are...Rich's! :P

	~ @venantius
*/

// NOTE: Implements IEditableCollection, IObj, IMapIterable, IKVReduce
type PersistentHashMap struct {
	APersistentMap

	_meta     IPersistentMap
	count     int
	root      INode
	hasNull   bool
	nullValue interface{}
}

var EMPTY_PERSISTENT_HASH_MAP = &PersistentHashMap{
	count:     0,
	root:      nil,
	hasNull:   false,
	nullValue: nil,
}

var NOT_FOUND interface{}
