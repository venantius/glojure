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

func CreatePersistentHashMapFromMap(other map[interface{}]interface{}) IPersistentMap {
	ret := EMPTY_PERSISTENT_HASH_MAP.AsTransient()
	for o := MapEntrySet(other).Seq(); o != nil; o = o.Next() {
		e := o.First().(MapEntry)
		ret = ret.Assoc(e.GetKey(), e.GetValue()).(*TransientHashMap)
	}
	return ret.Persistent().(*PersistentHashMap)
}

// TODO:
func CreatePersistentHashMap(init ...[]interface{}) *PersistentHashMap {
	return nil
}

// TODO
func CreatePersistentHashMapWithCheck(init interface{}) *PersistentHashMap {
	return nil
}

func hashPersistentHashMap(k interface{}) int {
	return Util.HashEq(k)
}

func (m *PersistentHashMap) ContainsKey(key interface{}) bool {
	if key == nil {
		return m.hasNull
	}
	if m.root != nil {
		return m.root.Find(0, hashPersistentHashMap(key), key, NOT_FOUND) != NOT_FOUND
	} else {
		return false
	}
}

func (m *PersistentHashMap) EntryAt(key interface{}) IMapEntry {
	if key == nil {
		if m.hasNull {
			return CreateMapEntry(nil, m.nullValue)
		} else {
			return nil
		}
	}
	if m.root != nil {
		return m.root.Find(0, hashPersistentHashMap(key), key, nil)
	} else {
		return nil
	}
}

// TODO
func (m *PersistentHashMap) Assoc(key interface{}, val interface{}) Associative {
	if key == nil {
		if m.hasNull && val == m.nullValue {
			return m
		}
		var c int
		if hasNull {
			c = m.count
		} else {
			c = m.count + 1
		}
		return &PersistentHashMap{
			meta:      m.Meta(),
			count:     c,
			root:      m.root,
			hasNull:   true,
			nullValue: val,
		}
	}
	addedLeaf := Box{}
	// TODO: More here, christ the use of ternary operators makes this verbose
	return nil
}

// TODO
func (m *PersistentHashMap) ValAt(key interface{}, notFound interface{}) interface{} {
	return nil
}

// TODO
func (m *PersistentHashMap) AssocEx(key interface{}, val interface{}) IPersistentMap {
	return nil
}

// TODO
func (m *PersistentHashMap) Without(key interface{}) IPersistentMap {
	return nil
}

// TODO: EMPTY_ITER is abstract iterator class here.

// TODO
func (m *PersistentHashMap) Iterator(f IFn) *Iterator {
	return nil
}

// TODO: Also, KeyIterator() and ValIterator() func

// TODO
func (m *PersistentHashMap) KVReduce(f IFn, init interface{}) interface{} {
	return nil
}

// TODO
func (m *PersistentHashMap) Fold(n int, combinef IFn, reducef IFn, fjinvoke IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	return nil
}

func (m *PersistentHashMap) Count() int {
	return m.count
}

func (m *PersistentHashMap) Seq() ISeq {
	var s ISeq
	if m.root != nil {
		s = m.root.NodeSeq()
	} else {
		s = nil
	}
	return s
}

func (m *PersistentHashMap) Empty() IPersistentCollection {
	return EMPTY_PERSISTENT_HASH_MAP.WithMeta(m.Meta())
}

// TODO
func Mask(hash int, shift int) int {
	return 0
}

func (m *PersistentHashMap) WithMeta(meta IPersistentMap) *PersistentHashMap {
	return &PersistentHashMap{
		_meta:     meta,
		count:     m.count,
		root:      m.root,
		hasNull:   m.hasNull,
		nullValue: m.nullValue,
	}
}

// TODO
func (m *PersistentHashMap) AsTransient() *TransientHashMap {
	return &TransientHashMap{
		_meta:     m._meta,
		edit:      true,
		root:      m.root,
		count:     m.count,
		hasNull:   m.hasNull,
		nullValue: m.nullValue,
	}
}

func (m *PersistentHashMap) Meta() IPersistentMap {
	return m._meta
}

// TODO: This.
type TransientHashMap struct {
	ATransientMap

	_meta     IPersistentMap
	edit      bool
	root      INode
	count     int
	hasNull   bool
	nullValue interface{}
	leafFlag  Box
}

// TODO
func (t *TransientHashMap) DoAssoc(key interface{}, val interface{}) ITransientMap {
	return nil
}

func (t *TransientHashMap) doPersistent() IPersistentCollection {
	t.edit = false
	return &PersistentHashMap{
		count:     t.count,
		root:      t.root,
		hasNull:   t.hasNull,
		nullValue: t.nullValue,
	}
}

// TODO: the rest of this file.
