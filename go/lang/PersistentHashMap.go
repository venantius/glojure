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

/*
	PersistentHashMap

	Implements: IEditableCollection, IObj, IMapIterable, IKVReduce
*/

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
		return m.root.Find(0, hashPersistentHashMap(key), key, nil).(IMapEntry)
	} else {
		return nil
	}
}

func (m *PersistentHashMap) Assoc(key interface{}, val interface{}) Associative {
	if key == nil {
		if m.hasNull && val == m.nullValue {
			return m
		}
		var c int
		if m.hasNull {
			c = m.count
		} else {
			c = m.count + 1
		}
		return &PersistentHashMap{
			_meta:     m.Meta(),
			count:     c,
			root:      m.root,
			hasNull:   true,
			nullValue: val,
		}
	}
	addedLeaf := Box{}
	var newroot INode
	if m.root == nil {
		newroot = EMPTY_BITMAP_INDEXED_NODE
	} else {
		newroot = m.root
	}
	newroot = newroot.Assoc(0, hashPersistentHashMap(key), key, val, addedLeaf)
	if newroot == m.root {
		return m
	}
	var c int
	if addedLeaf.val == nil {
		c = m.count
	} else {
		c = m.count + 1
	}
	return &PersistentHashMap{
		_meta:     m.Meta(),
		count:     c,
		root:      newroot,
		hasNull:   m.hasNull,
		nullValue: m.nullValue,
	}
}

func (m *PersistentHashMap) ValAt(key interface{}, notFound interface{}) interface{} {
	if key == nil {
		if m.hasNull {
			return m.hasNull
		} else {
			return notFound
		}
	}
	if m.root != nil {
		return m.root.Find(0, hashPersistentHashMap(key), key, notFound)
	} else {
		return notFound
	}
}

func (m *PersistentHashMap) AssocEx(key interface{}, val interface{}) IPersistentMap {
	if m.ContainsKey(key) {
		panic("Key already present")
	}
	return m.Assoc(key, val).(*PersistentHashMap)
}

func (m *PersistentHashMap) Without(key interface{}) IPersistentMap {
	if key == nil {
		if m.hasNull {
			return &PersistentHashMap{
				_meta:     m.Meta(),
				count:     m.count - 1,
				root:      m.root,
				hasNull:   false,
				nullValue: nil,
			}
		} else {
			return m
		}
	}
	if m.root == nil {
		return m
	}
	var newroot INode = m.root.Without(0, hashPersistentHashMap(key), key)
	if newroot == m.root {
		return m
	}
	return &PersistentHashMap{
		_meta:     m.Meta(),
		count:     m.count - 1,
		root:      newroot,
		hasNull:   m.hasNull,
		nullValue: m.nullValue,
	}
}

// TODO: EMPTY_ITER is abstract iterator class here.

// TODO
func (m *PersistentHashMap) Iterator(f IFn) *Iterator {
	return nil
}

// TODO: Also, KeyIterator() and ValIterator() func

func (m *PersistentHashMap) KVReduce(f IFn, init interface{}) interface{} {
	if m.hasNull {
		init = f.Invoke(init, nil, m.nullValue)
	}
	if RT.IsReduced(init) {
		return init.(IDeref).Deref()
	}
	if m.root != nil {
		init = m.root.KVReduce(f, init)
		if RT.IsReduced(init) {
			return init.(IDeref).Deref()
		} else {
			return init
		}
	}
	return init
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

func (m *PersistentHashMap) WithMeta(meta IPersistentMap) *PersistentHashMap {
	return &PersistentHashMap{
		_meta:     meta,
		count:     m.count,
		root:      m.root,
		hasNull:   m.hasNull,
		nullValue: m.nullValue,
	}
}

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

/*
	TransientHashMap
*/

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
func (t *TransientHashMap) doAssoc(key interface{}, val interface{}) ITransientMap {
	if key == nil {
		if t.nullValue != val {
			t.nullValue = val
		}
		if !t.hasNull {
			t.count++
			t.hasNull = true
		}
		return t
	}
	t.leafFlag.val = nil
	var n INode
	if t.root == nil {
		n = EMPTY_BITMAP_INDEXED_NODE
	} else {
		n = t.root
	}
	n.AssocWithEdit(t.edit, 0, hashPersistentHashMap(key), key, val, t.leafFlag)
	if n != t.root {
		t.root = n
	}
	if t.leafFlag.val != nil {
		t.count++
	}
	return t
}

// TODO
func (t *TransientHashMap) doWithout(key interface{}) ITransientMap {
	if key == nil {
		if !t.hasNull {
			return t
		}
		t.hasNull = false
		t.nullValue = nil
		t.count--
		return t
	}
	if t.root == nil {
		return t
	}
	t.leafFlag.val = nil
	var n INode = t.root.WithoutWithEdit(t.edit, 0, hashPersistentHashMap(key), key, t.leafFlag)
	if n != t.root {
		t.root = n
	}
	if t.leafFlag.val != nil {
		t.count--
	}
	return t
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

func (t *TransientHashMap) doValAt(key interface{}, notFound interface{}) interface{} {
	if key == nil {
		if t.hasNull {
			return t.nullValue
		} else {
			return notFound
		}
	}
	if t.root == nil {
		return notFound
	}
	return t.root.Find(0, hashPersistentHashMap(key), key, notFound)
}

func (t *TransientHashMap) doCount() int {
	return t.count
}

func (t *TransientHashMap) ensureEditable() {
	if t.edit == false {
		panic(TransientUsedAfterPersistentCallError)
	}
}

/*
	ArrayNode

	Implements: INode
*/

type ArrayNode struct {
	count int
	array []INode
	edit  bool
}

// TODO
func (n *ArrayNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	idx := Mask(hash, shift)
	node := n.array[idx]
	if node == nil {
		var arr []INode
		arr = cloneAndSet(
			n.array,
			idx,
			EMPTY_BITMAP_INDEXED_NODE.Assoc(
				shift+5,
				hash,
				key,
				val,
				addedLeaf,
			),
		)
		return &ArrayNode{
			edit:  false,
			count: n.count + 1,
			array: arr,
		}
	}
	var in INode = node.Assoc(shift+5, hash, key, val, addedLeaf)
	if in == node {
		return n
	}
	return &ArrayNode{
		edit:  false,
		count: n.count,
		array: cloneAndSet(n.array, idx, in),
	}
}

// TODO
func (n *ArrayNode) Without(shift int, hash int, key interface{}) INode {
	return nil
}

// TODO
func (n *ArrayNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	return nil
}

// TODO
func (n *ArrayNode) NodeSeq() ISeq {
	return nil
}

// TODO
func (n *ArrayNode) Iterator(f IFn) *Iterator {
	return nil
}

// TODO
func (n *ArrayNode) KVReduce(f IFn, init interface{}) interface{} {
	return nil
}

// TODO
func (n *ArrayNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	return nil
}

// TODO: FoldTasks

// TODO
func (n *ArrayNode) ensureEditable(edit bool) *ArrayNode {
	return nil
}

// TODO
func (n *ArrayNode) editAndSet(edit bool, i int, node INode) *ArrayNode {
	return nil
}

// TODO
func (n *ArrayNode) pack(edit bool, idx int) INode {
	return nil
}

// TODO
func (n *ArrayNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	return nil
}

// TODO
func (n *ArrayNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf Box) INode {
	return nil
}

/*
	ArrayNodeSeq
*/

type ArrayNodeSeq struct {
	ASeq

	_meta IPersistentMap
	nodes []INode
	i     int
	s     ISeq
}

func CreateArrayNodeSeq(meta IPersistentMap, nodes []INode, i int, s ISeq) ISeq {
	if s != nil {
		return &ArrayNodeSeq{
			_meta: meta,
			nodes: nodes,
			i:     i,
			s:     s,
		}
	}
	for j := i; j < len(nodes); j++ {
		if nodes[j] != nil {
			ns := nodes[j].NodeSeq()
			if ns != nil {
				return &ArrayNodeSeq{
					_meta: meta,
					nodes: nodes,
					i:     j + 1,
					s:     ns,
				}
			}
		}
	}
	return nil
}

func (s *ArrayNodeSeq) WithMeta(meta IPersistentMap) interface{} {
	return &ArrayNodeSeq{
		_meta: meta,
		nodes: s.nodes,
		i:     s.i,
		s:     s.s,
	}
}

func (s *ArrayNodeSeq) First() interface{} {
	return s.First()
}

func (s *ArrayNodeSeq) Next() ISeq {
	return CreateArrayNodeSeq(nil, s.nodes, s.i, s.s.Next())
}

/*
	ArrayNodeIter
*/

// TODO: This is an Iterator class.

/*
	BitmapIndexedNode
*/

// NOTE: Implements INode
type BitmapIndexedNode struct {
	bitmap int
	array  []interface{}
	edit   bool // TODO: Again, with the thread-locking. Not sure what the deal is here though
}

var EMPTY_BITMAP_INDEXED_NODE = &BitmapIndexedNode{
	edit:   false,
	bitmap: 0,
	array:  make([]interface{}, 0),
}

// TODO
func (n *BitmapIndexedNode) index() int {
	return 0
	// return Integer.bitCount(bitmap & (bit - 1))
}

// TODO
func (n *BitmapIndexedNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	return nil
}

// TODO
func (n *BitmapIndexedNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	return nil
}

// TODO
func (n *BitmapIndexedNode) Without(shift int, hash int, key interface{}) INode {
	return nil
}

// TODO
func (n *BitmapIndexedNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf Box) INode {
	return nil
}

// TODO
func (n *BitmapIndexedNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	return nil
}

// TODO
func (n *BitmapIndexedNode) NodeSeq() ISeq {
	return nil
}

// TODO
func (n *BitmapIndexedNode) Iterator(f IFn) *Iterator {
	return nil
}

// TODO
func (n *BitmapIndexedNode) KVReduce(f IFn, init interface{}) interface{} {
	return nil
}

// TODO
func (n *BitmapIndexedNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	return nil
}

// TODO
// NOTE: This also has another version that takes an AtomicReference to a thread
// as an argument.
func (n *BitmapIndexedNode) ensureEditable(edit bool) *BitmapIndexedNode {
	return nil
}

// TODO
// NOTE: overloaded
func (n *BitmapIndexedNode) editAndSet(edit bool, i int, a interface{}, j int, b interface{}) *BitmapIndexedNode {
	return nil
}

// TODO
func (n *BitmapIndexedNode) editAndRemovePair(edit bool, bit int, i int) *BitmapIndexedNode {
	return nil
}

/*
	NodeIter
*/

// NOTE: Implements Iterator
type NodeIter struct {
	array     []interface{}
	f         IFn
	i         int
	nextEntry interface{}
	nextIter  Iterator
}

// TODO
func (n *NodeIter) advance() bool {
	return true
}

// TODO
func (n *NodeIter) HasNext() bool {
	return true
}

// TODO
func (n *NodeIter) Next() interface{} {
	return nil
}

// TODO
func (n *NodeIter) Remove() {
	panic(UnsupportedOperationException)
}

/*
	NodeSeq
*/

type NodeSeq struct {
	ASeq

	_meta IPersistentMap
	array []interface{}
	i     int
	s     ISeq
}

// TODO
// NOTE: Overloaded
func CreateNodeSeq(array []interface{}) ISeq {
	return nil
}

// TODO
func KVReduceNodeSeq(array []interface{}, f IFn, init interface{}) interface{} {
	return nil
}

func (s *NodeSeq) WithMeta(meta IPersistentMap) *NodeSeq {
	return &NodeSeq{
		_meta: meta,
		array: s.array,
		i:     s.i,
		s:     s.s,
	}
}

func (s *NodeSeq) First() interface{} {
	if s.s != nil {
		return s.s.First()
	}
	return CreateMapEntry(s.array[s.i], s.array[s.i+1])
}

func (s *NodeSeq) Next() ISeq {
	if s.s != nil {
		return &NodeSeq{
			array: s.array,
			i:     s.i,
			s:     s.Next(),
		}
	}
	return &NodeSeq{
		array: s.array,
		i:     s.i + 2,
		s:     nil,
	}
}

/*
	HashCollisionNode
*/

// NOTE: Implements INode
type HashCollisionNode struct {
	hash  int
	count int
	array []interface{}
	edit  bool
}

// TODO
func (n *HashCollisionNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	return nil
}

// TODO
func (n *HashCollisionNode) Without(shift int, hash int, key interface{}) INode {
	return nil
}

// TODO
func (n *HashCollisionNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	return nil
}

// TODO
func (n *HashCollisionNode) NodeSeq() ISeq {
	return nil
}

// TODO
func (n *HashCollisionNode) Iterator(f IFn) *Iterator {
	return nil
}

// TODO
func (n *HashCollisionNode) KVReduce(f IFn, init interface{}) interface{} {
	return nil
}

// TODO
func (n *HashCollisionNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	return nil
}

// TODO
func (n *HashCollisionNode) FindIndex(key interface{}) int {
	return 0
}

// TODO
func (n *HashCollisionNode) ensureEditableNode(edit bool) *HashCollisionNode {
	return nil
}

// TODO
func (n *HashCollisionNode) ensureEditableWithArgs(edit bool, count int, array []interface{}) *HashCollisionNode {
	return nil
}

// TODO
func (n *HashCollisionNode) editAndSet(edit bool, i int, a interface{}, j int, b interface{}) *HashCollisionNode {
	return nil
}

// TODO
func (n *HashCollisionNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf Box) INode {
	return nil
}

// TODO
func (n *HashCollisionNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf Box) INode {
	return nil
}

/*
	Assorted static functions
*/

// NOTE: I think this is right but I'm not totally sure.
func Mask(hash int, shift int) uint {
	return (uint(hash) >> uint(shift)) & 31
}

func hashPersistentHashMap(k interface{}) int {
	return Util.HashEq(k)
}

// TODO
// NOTE: This is an overloaded method
func cloneAndSet(array []interface{}, i int, a interface{}) []interface{} {
	return nil
}

// cloneAndSetPair is just an overloaded version of cloneAndSet
func cloneAndSetPair(array []interface{}, i int, a interface{}, j int, b interface{}) []interface{} {
	return nil
}

func removePair(array []interface{}, i int) []interface{} {
	newArray := make([]interface{}, len(array)-2)
	copy(newArray[:2*i], array[:2*i])
	copy(newArray[2*(i+1):len(newArray)-2*i], array[2*i:len(newArray)-2*i])
	return newArray
}

// TODO
func createNode(edit bool, shift int, key1 interface{}, val1 interface{}, key2hash int, key2 interface{}, val2 interface{}) INode {
	return nil
}

func bitpos(hash int, shift int) int {
	return 1 << Mask(hash, shift)
}
