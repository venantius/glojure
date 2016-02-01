package lang
import "fmt"

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

	Implements abstract class: APersistentMap

	Implements: IEditableCollection, IObj, IMapIterable, IKVReduce
*/

type PersistentHashMap struct {
	AFn

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

func CreatePersistentHashMapFromSeq(items ISeq) *PersistentHashMap {
	ret := EMPTY_PERSISTENT_HASH_MAP.AsTransient()
	for ; items != nil; items = items.Next().Next() {
		ret = ret.Assoc(items.First(), RT.Second(items)).(*TransientHashMap)
	}
	return ret.Persistent().(*PersistentHashMap)
}

func CreatePersistentHashMapFromSeqWithCheck(init ...interface{}) *PersistentHashMap {
	ret := EMPTY_PERSISTENT_HASH_MAP.AsTransient()
	for i := 0 ; i < len(init); i += 2 {
		ret = ret.Assoc(init[i], init[i+1]).(*TransientHashMap)
	}
	return nil
}

func CreatePersistentHashMap(init ...interface{}) *PersistentHashMap {
	ret := EMPTY_PERSISTENT_HASH_MAP.AsTransient()
	for i := 0 ; i < len(init); i += 2 {
		ret = ret.Assoc(init[i], init[i+1]).(*TransientHashMap)
	}
	return ret.Persistent().(*PersistentHashMap)
}

func CreatePersistentHashMapWithCheck(init ...interface{}) *PersistentHashMap {
	var ret ITransientMap = EMPTY_PERSISTENT_HASH_MAP.AsTransient()
	for i := 0 ; i < len(init); i += 2 {
		ret = ret.Assoc(init[i], init[i+1])
		if ret.Count() != i/2 + 1 {
			panic(fmt.Sprintf("Duplicate key: %v", init[i]))
		}
	}
	return ret.Persistent().(*PersistentHashMap)
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
	addedLeaf := &Box{}
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
	panic(NotYetImplementedException)
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
	panic(NotYetImplementedException)
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
		leafFlag:  &Box{},
	}
}

func (m *PersistentHashMap) Meta() IPersistentMap {
	return m._meta
}

/*
	Abstract methods (PersistentHashMap)
 */

func (m *PersistentHashMap) Cons(o interface{}) IPersistentCollection {
	return APersistentMap_Cons(m, o)
}

func (m *PersistentHashMap) Equals(o interface{}) bool {
	return APersistentMap_Equals(m, o)
}

func (m *PersistentHashMap) Equiv(o interface{}) bool {
	return APersistentMap_Equiv(m, o)
}

func (m *PersistentHashMap) HashEq() int {
	return APersistentMap_HashEq(m)
}

/*
	TransientHashMap

	Implements abstract class ATransientMap
*/

type TransientHashMap struct {
	AFn

	_meta     IPersistentMap
	edit      bool
	root      INode
	count     int
	hasNull   bool
	nullValue interface{}
	leafFlag  *Box
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
	n = n.AssocWithEdit(t.edit, 0, hashPersistentHashMap(key), key, val, t.leafFlag)
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

func (t *TransientHashMap) doPersistent() IPersistentMap {
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
	Abstract methods (TransientHashMap)
 */

func (t *TransientHashMap) Assoc(key interface{}, val interface{}) ITransientMap {
	return ATransientMap_Assoc(t, key, val)
}

func (t *TransientHashMap) Conj(o interface{}) ITransientCollection {
	return ATransientMap_Conj(t, o)
}

func (t *TransientHashMap) Count() int {
	return ATransientMap_Count(t)
}

func (t *TransientHashMap) Persistent() IPersistentCollection {
	return ATransientMap_Persistent(t)
}

func (t *TransientHashMap) ValAt(key interface{}, notFound interface{}) interface{} {
	return ATransientMap_ValAt(t, key, notFound)
}

func (t *TransientHashMap) Without(key interface{}) ITransientMap {
	return ATransientMap_Without(t, key)
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
func (n *ArrayNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	idx := int(Mask(hash, shift))
	node := n.array[idx]
	if node == nil {
		var arr []INode
		arr = cloneAndSetINodeArray(
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
		array: cloneAndSetINodeArray(n.array, idx, in),
	}
}

// TODO
func (n *ArrayNode) Without(shift int, hash int, key interface{}) INode {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) NodeSeq() ISeq {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) Iterator(f IFn) *Iterator {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) KVReduce(f IFn, init interface{}) interface{} {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	panic(NotYetImplementedException)
}

// TODO: FoldTasks

// TODO
func (n *ArrayNode) ensureEditable(edit bool) *ArrayNode {
	panic(NotYetImplementedException)
}

func (n *ArrayNode) editAndSet(edit bool, i int, node INode) *ArrayNode {
	var editable *ArrayNode = n.ensureEditable(edit)
	editable.array[i] = n
	return editable
}

// TODO
func (n *ArrayNode) pack(edit bool, idx int) INode {
	panic(NotYetImplementedException)
}

// TODO
func (n *ArrayNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	var idx int = int(Mask(hash, shift))
	var node INode = n.array[idx]
	if node == nil {
		var editable *ArrayNode = n.editAndSet(edit, idx, EMPTY_BITMAP_INDEXED_NODE.AssocWithEdit(edit, shift+5, hash, key, val, addedLeaf))
		editable.count++
		return editable
	}
	var n2 INode = node.AssocWithEdit(edit, shift + 5, hash, key, val, addedLeaf)
	// TODO: Verify the following equality check here actually works.
	if n2 == node {
		return n2
	}
	return n.editAndSet(edit, idx, n2)
}

// TODO
func (n *ArrayNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf *Box) INode {
	panic(NotYetImplementedException)
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

func (n *BitmapIndexedNode) index(bit int) int {
	return Bitcount(n.bitmap & (bit - 1))
}

// TODO
func (n *BitmapIndexedNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	panic(NotYetImplementedException)
}

func (bin *BitmapIndexedNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	var bit int = bitpos(hash, shift)
	var idx int = bin.index(bit)
	if (bin.bitmap & bit) != 0 {
		keyOrNull := bin.array[2*idx]
		valOrNode := bin.array[2*idx+1]
		if keyOrNull == nil {
			var n INode = valOrNode.(INode).AssocWithEdit(edit, shift + 5, hash, key, val, addedLeaf)
			if n == valOrNode {
				return bin
			}
			return bin.editAndSetOne(edit, 2*idx+1, n)
		}
		if Util.Equiv(key, keyOrNull) {
			if val == valOrNode {
				return bin
			}
			return bin.editAndSetOne(edit, 2*idx+1, val)
		}
		addedLeaf.val = addedLeaf
		return bin.editAndSetTwo(edit, 2*idx, nil, 2*idx+1, createNode(edit, shift+5, keyOrNull, valOrNode, hash, key, val))
	} else {
		var n int = Bitcount(bin.bitmap)
		if (n*2 < len(bin.array)) {
			addedLeaf.val = addedLeaf
			var editable *BitmapIndexedNode = bin.ensureEditable(edit)
			copy(editable.array[2*(idx+1):2*(n-idx)], editable.array[2*idx:2*(n-idx)])
			editable.array[2*idx] = key
			editable.array[2*idx+1] = val
			editable.bitmap |= bit
			return editable
		}
		if n >= 16 {
			var nodes []INode = make([]INode, 32)
			var jdx uint = Mask(hash, shift)
			nodes[jdx] = EMPTY_BITMAP_INDEXED_NODE.AssocWithEdit(edit, shift + 5, hash, key, val, addedLeaf)
			var j int = 0
			for i := uint(0); i < 32; i++ {
				if ((bin.bitmap >> i) & 1) != 0 {
					if bin.array[j] == nil {
						nodes[i] = bin.array[j+1].(INode)
					} else {
						nodes[i] = EMPTY_BITMAP_INDEXED_NODE.AssocWithEdit(edit, shift+5, hashPersistentHashMap(bin.array[j]), bin.array[j], bin.array[j+1], addedLeaf)
					}
					j += 2
				}
			}
			return &ArrayNode{
				edit: edit,
				count: n + 1,
				array: nodes,
			}
		} else {
			var newArray []interface{} = make([]interface{}, 2*(n+4))
			copy(newArray[:2*idx], bin.array[:2*idx])
			newArray[2*idx] = key
			addedLeaf.val = addedLeaf
			newArray[2*idx+1] = val
			copy(newArray[2*(idx+1):2*(idx+1)+2*(n-idx)], bin.array[2*idx:2*idx+2*(n-idx)])
			var editable *BitmapIndexedNode = bin.ensureEditable(edit)
			editable.array = newArray
			editable.bitmap |= bit
			return editable
		}
	}
}

// TODO
func (n *BitmapIndexedNode) Without(shift int, hash int, key interface{}) INode {
	panic(NotYetImplementedException)
}

// TODO
func (n *BitmapIndexedNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf *Box) INode {
	panic(NotYetImplementedException)
}

func (n *BitmapIndexedNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	var bit int = bitpos(hash, shift)
	if (n.bitmap & bit) == 0 {
		return nil
	}
	var idx int = n.index(bit)
	keyOrNull := n.array[2*idx]
	valOrNode := n.array[2*idx+1]
	if keyOrNull == nil {
		return valOrNode.(INode).Find(shift+5, hash, key, notFound)
	}
	if Util.Equiv(key, keyOrNull) {
		return CreateMapEntry(keyOrNull, valOrNode)
	}
	return notFound
}

// TODO
func (n *BitmapIndexedNode) NodeSeq() ISeq {
	return CreateNodeSeq(n.array)
}

// TODO
func (n *BitmapIndexedNode) Iterator(f IFn) *Iterator {
	panic(NotYetImplementedException)
}

// TODO
func (n *BitmapIndexedNode) KVReduce(f IFn, init interface{}) interface{} {
	panic(NotYetImplementedException)
}

// TODO
func (n *BitmapIndexedNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	panic(NotYetImplementedException)
}

// NOTE: This also has another version that takes an AtomicReference to a thread
// as an argument.
// NOTE: See TransientVector.ensureEditable
func (bin *BitmapIndexedNode) ensureEditable(edit bool) *BitmapIndexedNode {
	if edit == true {
		return bin
	}
	var n int = Bitcount(bin.bitmap)
	var l int
	if n >= 0 {
		l = 2 * (n + 1)
	} else {
		l = 4
	}
	var newArray []interface{} = make([]interface{}, l)
	copy(newArray[:2*n], bin.array[:2*n])
	return &BitmapIndexedNode{
		edit: edit,
		bitmap: bin.bitmap,
		array: newArray,
	}
}

// NOTE: How did I do this in other circumstances?
func (n *BitmapIndexedNode) editAndSetOne(edit bool, i int, a interface{}) *BitmapIndexedNode {
	var editable *BitmapIndexedNode = n.ensureEditable(edit)
	editable.array[i] = a
	return editable
}

func (n *BitmapIndexedNode) editAndSetTwo(edit bool, i int, a interface{}, j int, b interface{}) *BitmapIndexedNode {
	var editable *BitmapIndexedNode = n.ensureEditable(edit)
	editable.array[i] = a
	editable.array[j] = b
	return editable
}

// TODO
func (n *BitmapIndexedNode) editAndRemovePair(edit bool, bit int, i int) *BitmapIndexedNode {
	panic(NotYetImplementedException)
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
	panic(NotYetImplementedException)
}

// TODO
func (n *NodeIter) HasNext() bool {
	panic(NotYetImplementedException)
}

// TODO
func (n *NodeIter) Next() interface{} {
	panic(NotYetImplementedException)
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
	return createNodeSeq(array, 0, nil)
}

func createNodeSeq(array []interface{}, i int, s ISeq) ISeq {
	if s != nil {
		return &NodeSeq{
			_meta: nil,
			array: array,
			i: i,
			s: s,
		}
	}
	for j := i; j < len(array); j += 2 {
		if array[j] != nil {
			return &NodeSeq{
				_meta: nil,
				array: array,
				i: j,
				s: nil,
			}
		}
		node := array[j+1]
		if node != nil {
			node := node.(INode)
			var nodeSeq ISeq = node.NodeSeq()
			if nodeSeq != nil {
				return &NodeSeq{
					_meta: nil,
					array: array,
					i: j + 2,
					s: nodeSeq,
				}
			}
		}
	}
	return nil
}

// TODO
func KVReduceNodeSeq(array []interface{}, f IFn, init interface{}) interface{} {
	panic(NotYetImplementedException)
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
		return createNodeSeq(s.array, s.i, s.s.Next())
	}
	return createNodeSeq(s.array, s.i + 2, nil)
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
func (n *HashCollisionNode) Assoc(shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	panic(NotYetImplementedException)
}

// TODO
func (n *HashCollisionNode) Without(shift int, hash int, key interface{}) INode {
	panic(NotYetImplementedException)
}

// TODO
func (n *HashCollisionNode) Find(shift int, hash int, key interface{}, notFound interface{}) interface{} {
	var idx int = n.FindIndex(key)
	if idx < 0 {
		return notFound
	}
	if Util.Equiv(key, n.array[idx]) {
		return n.array[idx+1]
	}
	return notFound
}

func (n *HashCollisionNode) NodeSeq() ISeq {
	return CreateNodeSeq(n.array)
}

// TODO
func (n *HashCollisionNode) Iterator(f IFn) *Iterator {
	panic(NotYetImplementedException)
}

// TODO
func (n *HashCollisionNode) KVReduce(f IFn, init interface{}) interface{} {
	panic(NotYetImplementedException)
}

// TODO
func (n *HashCollisionNode) Fold(combinef IFn, reducef IFn, fjtask IFn, fjfork IFn, fjjoin IFn) interface{} {
	panic(NotYetImplementedException)
}

func (n *HashCollisionNode) FindIndex(key interface{}) int {
	for i := 0; i < 2 * n.count; i += 2 {
		if Util.Equiv(key, n.array[i]) {
			fmt.Println("They were equals!", key, n.array[i])

			return i
		}
	}
	return -1
}

func (n *HashCollisionNode) ensureEditableNode(edit bool) *HashCollisionNode {
	if edit == true {
		return n
	}
	var newArray []interface{} = make([]interface{}, 2*(n.count+1))
	copy(newArray[:2*n.count], n.array[:2*n.count])
	return &HashCollisionNode{
		edit: edit,
		hash: n.hash,
		count: n.count,
		array: newArray,
	}
}

func (n *HashCollisionNode) ensureEditableWithArgs(edit bool, count int, array []interface{}) *HashCollisionNode {
	if edit == true {
		n.array = array
		n.count = count
		return n
	}
	return &HashCollisionNode{
		edit: edit,
		hash: n.hash,
		count: count,
		array: array,
	}
}

func (n *HashCollisionNode) editAndSetOne(edit bool, i int, a interface{}) *HashCollisionNode {
	var editable *HashCollisionNode = n.ensureEditableNode(edit)
	editable.array[i] = a
	return editable
}

func (n *HashCollisionNode) editAndSetTwo(edit bool, i int, a interface{}, j int, b interface{}) *HashCollisionNode {
	var editable *HashCollisionNode = n.ensureEditableNode(edit)
	editable.array[i] = a
	editable.array[j] = b
	return editable
}

// TODO
func (n *HashCollisionNode) AssocWithEdit(edit bool, shift int, hash int, key interface{}, val interface{}, addedLeaf *Box) INode {
	if hash == n.hash {
		var idx int = n.FindIndex(key)
		if idx != -1 {
			if n.array[idx+1] == val {
				return n
			}
			return n.editAndSetOne(edit, idx+1, val)
		}
		if len(n.array) > 2 * n.count {
			addedLeaf.val = addedLeaf
			var editable *HashCollisionNode = n.editAndSetTwo(edit, 2*n.count, key, 2*n.count+1, val)
			editable.count++
			return editable
		}
		var newArray []interface{} = make([]interface{}, len(n.array)+2)
		copy(newArray[:len(n.array)], n.array[:len(n.array)])
		newArray[len(n.array)] = key
		newArray[len(n.array)+1] = val
		addedLeaf.val = addedLeaf
		return n.ensureEditableWithArgs(edit, n.count+1, newArray)
	}
	var b *BitmapIndexedNode = &BitmapIndexedNode{
		edit: edit,
		bitmap: bitpos(n.hash, shift),
		array: make([]interface{}, 4),
	}
	return b.AssocWithEdit(edit, shift, hash, key, val, addedLeaf)
}

// TODO
func (n *HashCollisionNode) WithoutWithEdit(edit bool, shift int, hash int, key interface{}, removedLeaf *Box) INode {
	panic(NotYetImplementedException)
}

/*
	Assorted static functions
*/

// NOTE: I think this is right but I'm not totally sure.
func Mask(hash int, shift int) uint {
	return (uint(hash) >> uint(shift)) & 0x01f
}

func hashPersistentHashMap(k interface{}) int {
	return Util.HashEq(k)
}

// Another overloaded version of cloneAndSetINodeArray
func cloneAndSetINodeArray(array []INode, i int, a INode) []INode {
	var clone []INode
	copy(clone, array)
	clone[i] = a
	return clone
}

// NOTE: This is an overloaded method.
func cloneAndSet(array []interface{}, i int, a interface{}) []interface{} {
	var clone []interface{}
	copy(clone, array)
	clone[i] = a
	return clone
}

// cloneAndSetPair is just an overloaded version of cloneAndSet
func cloneAndSetPair(array []interface{}, i int, a interface{}, j int, b interface{}) []interface{} {
	var clone []interface{}
	copy(clone, array)
	clone[i] = a
	clone[j] = b
	return clone
}

func removePair(array []interface{}, i int) []interface{} {
	newArray := make([]interface{}, len(array)-2)
	copy(newArray[:2*i], array[:2*i])
	copy(newArray[2*(i+1):len(newArray)-2*i], array[2*i:len(newArray)-2*i])
	return newArray
}

func createNode(edit bool, shift int, key1 interface{}, val1 interface{}, key2hash int, key2 interface{}, val2 interface{}) INode {
	var key1hash int = hashPersistentHashMap(key1)
	if key1hash == key2hash {
		arr := make([]interface{}, 4)
		arr[0] = key1
		arr[1] = val1
		arr[2] = key2
		arr[3] = val2
		return &HashCollisionNode{
			edit: true, // TODO: Is this correct? Should this be initalized as false?
			hash: key1hash,
			count: 2,
			array: arr,
		}
	}
	var addedLeaf *Box = &Box{}
	return EMPTY_BITMAP_INDEXED_NODE.AssocWithEdit(edit, shift, key1hash, key1, val1, addedLeaf).AssocWithEdit(edit, shift, key2hash, key2, val2, addedLeaf)
}

func bitpos(hash int, shift int) int {
	return 1 << Mask(hash, shift)
}
