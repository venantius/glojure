package lang

import (
	"errors"
	"fmt" // really just for debugging
)

var emptyVectorPopError = "Can't pop empty vector."

// NOTE: Implements IObj, IEditableCollection, IReduce, IKVReduce
type PersistentVector struct {
	*APersistentVector

	cnt   int // count
	shift uint
	root  *Node
	tail  []interface{}
	_meta IPersistentMap
}

// NOTE: Implements Serializable
type Node struct {
	/*
		NOTE: In Clojure, `edit` is an AtomicReference<Thread>. Since Clojure 1.7.0
		transients have not had their thread-local checks enforced, however, and
		Go lacks an interface for managing threads directly anyways. Transients do
		still use the edit field to check for whether a `persistent` call has been
		made, and so we use a simple boolean to capture that here.
	*/
	edit  bool
	array []interface{}
}

const (
	VECTOR_SHIFT = 5
	NODE_SIZE    = 32
	NODE_END_IDX = NODE_SIZE - 1
)

var EMPTY_PERSISTENT_VECTOR_NODE = &Node{
	edit:  false,
	array: make([]interface{}, NODE_SIZE),
}

var EMPTY_PERSISTENT_VECTOR = &PersistentVector{
	cnt:   0,
	shift: VECTOR_SHIFT,
	root:  EMPTY_PERSISTENT_VECTOR_NODE,
	tail:  make([]interface{}, 0),
}

/*
	NOTE: This is just a small anonymous class in Java
*/
type transientVectorConj struct {
	AFn
	IFn // maybe?
}

// Original method is overloaded.
func (t *transientVectorConj) Invoke(args ...interface{}) interface{} {
	coll := args[0]
	val := args[1]
	if val == nil {
		return coll
	}
	return coll.(ITransientVector).Conj(val)
}

var TRANSIENT_VECTOR_CONJ = transientVectorConj{}

// Return a new PersistentVector with the items passed in.
func adopt(items []interface{}) *PersistentVector {
	return &PersistentVector{
		cnt:   len(items),
		shift: VECTOR_SHIFT,
		root:  EMPTY_PERSISTENT_VECTOR_NODE,
		tail:  items,
	}
}

func createVectorFromIReduceInit(items IReduceInit) *PersistentVector {
	ret := EMPTY_PERSISTENT_VECTOR.AsTransient()
	items.ReduceWithInit(&TRANSIENT_VECTOR_CONJ, ret)
	return ret.Persistent()
}

func createVectorFromISeq(items ISeq) *PersistentVector {
	arr := make([]interface{}, NODE_SIZE)
	i := 0
	for ; items != nil && i < NODE_SIZE; items = items.Next() {
		arr[i] = items.First()
		i++
	}
	if items != nil {
		start := PersistentVector{
			cnt:   NODE_SIZE,
			root:  EMPTY_PERSISTENT_VECTOR_NODE,
			shift: VECTOR_SHIFT,
			tail:  arr,
		}
		ret := start.AsTransient()
		for ; items != nil; items = items.Next() {
			ret = ret.Conj(items.First())
		}
		return ret.Persistent()
	} else if i == 32 {
		return &PersistentVector{
			cnt:   NODE_SIZE,
			root:  EMPTY_PERSISTENT_VECTOR_NODE,
			shift: VECTOR_SHIFT,
			tail:  arr,
		}
	} else {
		arr2 := make([]interface{}, i)
		copy(arr2, arr)
		return &PersistentVector{
			cnt:   i,
			shift: VECTOR_SHIFT,
			root:  EMPTY_PERSISTENT_VECTOR_NODE,
			tail:  arr2,
		}
	}
}

// TODO: This is complicated because it's tied to java.util.List, which is an
// abstraction that doesn't seem to exist in Go
func createVectorFromList(list *List) *PersistentVector {
	size := list.Size()
	if size <= NODE_SIZE {
		return &PersistentVector{
			cnt:   size,
			shift: VECTOR_SHIFT,
			root:  EMPTY_PERSISTENT_VECTOR_NODE,
			tail:  list.ToArray(),
		}
	}
	ret := EMPTY_PERSISTENT_VECTOR.AsTransient()
	for i := 0; i < size; i++ {
		ret = ret.Conj(list.Get(i))
	}
	return ret.Persistent()
}

// TODO: Review idiomatic iterators in Go, don't try to copy this from Java.
func createVectorFromIterable(items Iterable) *PersistentVector {
	// TODO: if arraylist, use createVectorFromList
	// iter := items.Iterator()
	ret := EMPTY_PERSISTENT_VECTOR.AsTransient()
	// TODO: this should be a while loop.
	// for iter.HasNext() {
	//	ret = ret.Conj(iter.Next())
	//}
	return ret.Persistent()
}

func createVectorFromInterfaceSlice(items []interface{}) *PersistentVector {
	ret := EMPTY_PERSISTENT_VECTOR.AsTransient()
	for _, item := range items {
		ret = ret.Conj(item)
	}
	return ret.Persistent()
}

// General initializer; does type checking and dispatches to appropriate
// constructor.
func CreateVector(items ...interface{}) *PersistentVector {
	ret := EMPTY_PERSISTENT_VECTOR
	switch item := items[0].(type) {
	case IReduceInit:
		ret = createVectorFromIReduceInit(item)
	case ISeq:
		ret = createVectorFromISeq(item)
	// TODO: uncomment me once we understand iterables in Go.
	// case Iterable:
	// 	ret = createVectorFromIterable(items[0].(Iterable))
	default:
		ret = createVectorFromInterfaceSlice(items)
	}
	return ret
}

func (v *PersistentVector) AsTransient() *TransientVector {
	return &TransientVector{
		cnt:   v.cnt,
		root:  editableRoot(v.root),
		shift: v.shift,
		tail:  editableTail(v.tail),
	}
}

func (v *PersistentVector) tailoff() int {
	if v.cnt < NODE_SIZE {
		return 0
	}
	// TODO: I'm not totally clear on whether this is consistent with the Java version.
	return ((v.cnt - 1) >> VECTOR_SHIFT) << VECTOR_SHIFT
}

func (v *PersistentVector) ArrayFor(i int) []interface{} {
	if i < 0 || i >= v.cnt {
		panic(IndexOutOfBoundsException)
	}
	if i >= v.tailoff() {
		return v.tail
	}
	n := v.root
	for level := v.shift; level > 0; level -= VECTOR_SHIFT {
		// NOTE: bitshift is probably wrong here as well.
		n = n.array[(i>>level)&(NODE_END_IDX)].(*Node)
	}
	return n.array
}

// Return the number of items in the vector.
func (v *PersistentVector) Count() int {
	return v.cnt
}

func (v *PersistentVector) nth(i int) interface{} {
	subsl := v.ArrayFor(i)
	return subsl[i&(NODE_END_IDX)]
}

// Retrieve the nth item in the vector. If the index being retrieved is beyond
// the length of the vector, returns the notFound value.
func (v *PersistentVector) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < v.cnt {
		return v.nth(i)
	}
	return notFound
}

// Render string representation. Allows for custom printing of PersistentVectors.
// TODO: This currently prints strings without quotation marks.
func (v PersistentVector) String() string {
	s := "["
	for i := 0; i < v.Count(); i++ {
		if i > 0 {
			s += " "
		}
		s += fmt.Sprint(v.Nth(i, nil))
	}
	s += "]"
	return s
}

// Assoc in a new value at the index.
func (v *PersistentVector) AssocN(i int, val interface{}) *PersistentVector {
	if i >= 0 && i < v.cnt {
		if i >= v.tailoff() {
			newTail := make([]interface{}, len(v.tail))
			copy(newTail, v.tail)
			newTail[i&(NODE_END_IDX)] = val
			return &PersistentVector{
				_meta: v.Meta(),
				cnt:   v.cnt,
				shift: v.shift,
				root:  v.root,
				tail:  newTail,
			}
		}
		return &PersistentVector{
			_meta: v.Meta(),
			cnt:   v.cnt,
			shift: v.shift,
			root:  doAssoc(v.shift, v.root, i, val),
			tail:  v.tail,
		}
	}
	if i == v.cnt {
		return v.Cons(val)
	}
	panic(IndexOutOfBoundsException)
}

// Private function to handle assoc-ing at a lower level
func doAssoc(level uint, node *Node, i int, val interface{}) *Node {
	var arr []interface{}
	copy(arr, node.array)
	ret := &Node{edit: node.edit, array: arr}
	if level == 0 {
		ret.array[i&(NODE_END_IDX)] = val
	} else {
		// NOTE: Bitwise issues again.
		subidx := (i >> level) & (NODE_END_IDX)
		ret.array[subidx] = doAssoc(level-VECTOR_SHIFT, node.array[subidx].(*Node), i, val)
	}
	return ret
}

// Return a new PersistentVector with new metadata.
func (v *PersistentVector) WithMeta(meta IPersistentMap) *PersistentVector {
	return &PersistentVector{
		_meta: meta,
		cnt:   v.cnt,
		shift: v.shift,
		root:  v.root,
		tail:  v.tail,
	}
}

// Return the PersistentVector's metadata.
func (v *PersistentVector) Meta() IPersistentMap {
	return v._meta
}

func newPath(edit bool, level uint, node *Node) *Node {
	if level == 0 {
		return node
	}
	ret := &Node{edit: edit}
	ret.array[0] = newPath(edit, level-VECTOR_SHIFT, node)
	return ret
}

func (v *PersistentVector) pushTail(level uint, parent *Node, tailnode *Node) *Node {
	// NOTE: bitshifts require review
	subidx := ((v.cnt - 1) >> level) & NODE_SIZE
	ret := &Node{edit: parent.edit, array: *&parent.array}
	nodeToInsert := &Node{edit: false}
	if level == VECTOR_SHIFT {
		nodeToInsert = tailnode
	} else {
		child := parent.array[subidx]
		if child != nil {
			nodeToInsert = v.pushTail(level-VECTOR_SHIFT, child.(*Node), tailnode)
		} else {
			nodeToInsert = newPath(v.root.edit, level-VECTOR_SHIFT, tailnode)
		}
	}
	ret.array[subidx] = nodeToInsert
	return ret
}

func (v *PersistentVector) Cons(val interface{}) *PersistentVector {
	if v.cnt-v.tailoff() < NODE_SIZE {
		newTail := make([]interface{}, len(v.tail)+1)
		copy(newTail, v.tail)
		newTail[len(v.tail)] = val
		return &PersistentVector{
			_meta: v.Meta(),
			cnt:   v.cnt + 1,
			shift: v.shift,
			root:  v.root,
			tail:  newTail,
		}
	}
	newroot := &Node{edit: false}
	tailnode := &Node{v.root.edit, v.tail}
	newshift := v.shift
	fmt.Println(tailnode)
	// NOTE: Again, not comfortable with bit shifting here.
	if (v.cnt >> VECTOR_SHIFT) > (1 << v.shift) {
		newroot = &Node{edit: v.root.edit} // defaults?
		newroot.array[0] = v.root
		newroot.array[1] = newPath(v.root.edit, v.shift, tailnode)
		newshift += VECTOR_SHIFT
	} else {
		newroot = v.pushTail(v.shift, v.root, tailnode)
	}
	return &PersistentVector{
		_meta: v.Meta(),
		cnt:   v.cnt + 1,
		shift: newshift,
		root:  newroot,
		tail:  []interface{}{val}}
}

func (v *PersistentVector) ChunkedSeq() IChunkedSeq {
	if v.Count() == 0 {
		return nil
	}
	return &ChunkedSeq{
		vec:    *v,
		i:      0,
		offset: 0,
	}
}

func (v *PersistentVector) Seq() ISeq {
	return v.ChunkedSeq()
}

// TODO: This will be hard
func (v *PersistentVector) rangedIterator(start int, end int) Iterator {
	return Iterator{}
}

func (v *PersistentVector) Iterator() Iterator {
	return v.rangedIterator(0, v.Count())
}

func (v *PersistentVector) Reduce(f IFn, init *interface{}) interface{} {
	// Handle the method overloading
	// NOTE: This method's overloading is somewhat complex, so we'll be careful.
	if init == nil {
		if v.cnt > 0 {
			_temp := v.ArrayFor(0)[0]
			init = &_temp
		} else {
			return f.Invoke()
		}
	}

	step := 0
	for i := 0; i < v.cnt; i += step {
		array := v.ArrayFor(i)
		// In Clojure these are pre-incremented
		for j := 0; j < len(array); j++ {
			init := f.Invoke(init, array[j])
			if RT.IsReduced(init) {
				return init.(IDeref).Deref()
			}
		}
		step = len(array)
	}
	return init
}

func (v *PersistentVector) KVReduce(f IFn, init interface{}) interface{} {
	step := 0
	for i := 0; i < v.cnt; i += step {
		array := v.ArrayFor(i)
		for j := 0; j < len(array); j++ {
			init := f.Invoke(init, j+i, array[j])
			if RT.IsReduced(init) {
				return init.(IDeref).Deref()
			}
		}
		step = len(array)
	}
	return init
}

// NOTE: implements IChunkedSeq, Counted
type ChunkedSeq struct {
	*ASeq

	vec    PersistentVector
	node   []interface{}
	i      int
	offset int
}

func (c *ChunkedSeq) ChunkedFirst() IChunk {
	return &ArrayChunk{
		array: c.node,
		off:   c.offset,
	}
}

func (c *ChunkedSeq) ChunkedNext() ISeq {
	if c.i+len(c.node) < c.vec.cnt {
		return &ChunkedSeq{vec: c.vec, i: c.i + len(c.node), offset: 0}
	}
	return nil
}

func (c *ChunkedSeq) ChunkedMore() ISeq {
	s := c.ChunkedNext()
	if s == nil {
		// TODO: This could probably be replaced with an EmptyList struct.
		return EMPTY_PERSISTENT_LIST
	}
	return s
}

func (c *ChunkedSeq) WithMeta(meta IPersistentMap) *ChunkedSeq {
	if meta == c.vec._meta {
		return c
	}
	return &ChunkedSeq{
		vec:    *c.vec.WithMeta(meta),
		node:   c.node,
		i:      c.i,
		offset: c.offset,
	}
}

func (c *ChunkedSeq) First() interface{} {
	return c.node[c.offset]
}

func (c *ChunkedSeq) Next() ISeq {
	if c.offset+1 < len(c.node) {
		return &ChunkedSeq{
			vec:    c.vec,
			node:   c.node,
			i:      c.i,
			offset: c.offset + 1,
		}
	}
	return c.ChunkedNext()
}

// Return the size of the chunked sequence.
func (c *ChunkedSeq) Count() int {
	return c.vec.cnt - (c.i + c.offset)
}

// Empty the vector's contents.
func (v *PersistentVector) Empty() *PersistentVector {
	return EMPTY_PERSISTENT_VECTOR.WithMeta(v.Meta())
}

func (v *PersistentVector) Pop() *PersistentVector {
	if v.cnt == 0 {
		panic(emptyVectorPopError)
	}
	if v.cnt == 1 {
		return EMPTY_PERSISTENT_VECTOR.WithMeta(v.Meta())
	}
	if (v.cnt - v.tailoff()) > 1 {
		newTail := make([]interface{}, len(v.tail)-1)
		copy(newTail, v.tail)
		return &PersistentVector{
			_meta: v.Meta(),
			cnt:   v.cnt - 1,
			shift: v.shift,
			root:  v.root,
			tail:  newTail,
		}
	}
	newTail := v.ArrayFor(v.cnt - 2)
	newroot := v.popTail(v.shift, v.root)
	newshift := v.shift
	if newroot == nil {
		newroot = EMPTY_PERSISTENT_VECTOR_NODE
	}
	if v.shift > VECTOR_SHIFT && &newroot.array[1] == nil {
		newroot = newroot.array[0].(*Node)
		newshift -= VECTOR_SHIFT
	}
	return &PersistentVector{
		_meta: v.Meta(),
		cnt:   v.cnt - 1,
		shift: newshift,
		root:  newroot,
		tail:  newTail,
	}
}

func (v *PersistentVector) popTail(level uint, node *Node) *Node {
	subidx := ((v.cnt - 2) >> level) & (NODE_END_IDX)
	if level > VECTOR_SHIFT {
		newchild := v.popTail(level-VECTOR_SHIFT, node.array[subidx].(*Node))
		if &newchild == nil && subidx == 0 {
			return nil
		} else {
			var arr []interface{}
			copy(arr, node.array)
			ret := Node{edit: v.root.edit, array: arr}
			ret.array[subidx] = newchild
			return &ret
		}
	} else if subidx == 0 {
		return nil
	} else {
		var arr []interface{}
		copy(arr, node.array)
		ret := Node{edit: v.root.edit, array: arr}
		ret.array[subidx] = nil
		return &ret
	}
}

// NOTE: Implements ITransientVector, Counted
type TransientVector struct {
	*AFn

	cnt   int
	shift uint
	root  *Node
	tail  []interface{}
}

func (t *TransientVector) Count() int {
	t.ensureEditable()
	return t.cnt
}

func (t *TransientVector) ensureEditable() {
	// NOTE: t.root.edit.get(), atomically in Java
	if t.root.edit == false {
		panic("Transient used after persistent! call")
	}
}

func (t *TransientVector) ensureEditableNode(node *Node) *Node {
	if node.edit == t.root.edit {
		return node
	}
	var arr []interface{}
	copy(arr, node.array)
	return &Node{edit: t.root.edit, array: arr}
}

func editableRoot(node *Node) *Node {
	var arr []interface{}
	copy(arr, node.array)
	return &Node{
		edit:  true,
		array: arr,
	}
}

func (t *TransientVector) Persistent() *PersistentVector {
	t.ensureEditable()
	t.root.edit = false
	trimmedTail := make([]interface{}, t.cnt-t.tailoff())
	copy(trimmedTail, t.tail)
	return &PersistentVector{
		cnt:   t.cnt,
		shift: t.shift,
		root:  t.root,
		tail:  trimmedTail,
	}
}

func editableTail(t []interface{}) []interface{} {
	arr := make([]interface{}, NODE_SIZE)
	copy(arr, t)
	return arr
}

func (t *TransientVector) Conj(val interface{}) *TransientVector {
	t.ensureEditable()
	i := t.cnt
	if (i - t.tailoff()) < NODE_SIZE {
		t.tail[i&(NODE_END_IDX)] = val
		t.cnt++
		return t
	}
	var newroot *Node
	tailnode := &Node{
		edit:  t.root.edit,
		array: t.tail,
	}
	tail := make([]interface{}, NODE_SIZE)
	tail[0] = val
	newshift := t.shift
	// TODO: review bit shift
	if (t.cnt >> VECTOR_SHIFT) > (1 << t.shift) {
		newroot = &Node{edit: t.root.edit}
		newroot.array[0] = t.root
		newroot.array[1] = newPath(t.root.edit, t.shift, tailnode)
		newshift += VECTOR_SHIFT
	} else {
		newroot = t.pushTail(t.shift, t.root, tailnode)
	}
	t.root = newroot
	t.shift = newshift
	t.cnt++
	return t
}

func (t *TransientVector) pushTail(level uint, parent *Node, tailnode *Node) *Node {
	parent = t.ensureEditableNode(parent)
	// TODO: bit shifting?
	subidx := ((t.cnt - 1) >> level) & (NODE_END_IDX)
	ret := parent
	var nodeToInsert *Node
	if level == VECTOR_SHIFT {
		nodeToInsert = tailnode
	} else {
		child := parent.array[subidx]
		if child != nil {
			nodeToInsert = t.pushTail(level-VECTOR_SHIFT, child.(*Node), tailnode)
		} else {
			nodeToInsert = newPath(t.root.edit, level-VECTOR_SHIFT, tailnode)
		}
	}
	ret.array[subidx] = nodeToInsert
	return ret
}

func (t *TransientVector) tailoff() int {
	if t.cnt < NODE_SIZE {
		return 0
	}
	// TODO: Bitshifts
	return ((t.cnt - 1) >> VECTOR_SHIFT) << VECTOR_SHIFT
}

func (t *TransientVector) arrayFor(i int) []interface{} {
	if i >= 0 && i < t.cnt {
		if i >= t.tailoff() {
			return t.tail
		}
		node := t.root
		for level := t.shift; level > 0; level -= VECTOR_SHIFT {
			// TODO: bitshift
			node = node.array[(i>>level)&(NODE_END_IDX)].(*Node)
		}
		return node.array
	}
	panic(IndexOutOfBoundsException)
}

func (t *TransientVector) editableArrayFor(i int) []interface{} {
	if i >= 0 && i < t.cnt {
		if i >= t.tailoff() {
			return t.tail
		}
		node := t.root
		for level := t.shift; level > 0; level -= VECTOR_SHIFT {
			// TODO: bit shift
			node = t.ensureEditableNode(node.array[(i>>level)&(NODE_END_IDX)].(*Node))
		}
		return node.array
	}
	panic(IndexOutOfBoundsException)
}

// NOTE: Function overloading
// Retrieve the value at the corresponding index of this TransientVector
func (t *TransientVector) ValAt(key interface{}, notFound interface{}) interface{} {
	t.ensureEditable()
	switch k := key.(type) {
	case int:
		i := k
		if i >= 0 && i < t.cnt {
			return t.Nth(i, nil)
		}
	}
	return notFound
}

func (t *TransientVector) Invoke(arg1 interface{}) interface{} {
	switch a := arg1.(type) {
	case int:
		return t.Nth(a, nil)
	default:
		panic(errors.New("Key must be integer"))
	}
}

// NOTE: Function overloaded in Java
// Return the value at the Nth index of the TransientVector.
func (t *TransientVector) Nth(i int, notFound interface{}) interface{} {
	t.ensureEditable()
	node := t.arrayFor(i)
	if i >= 0 && i < t.Count() {
		return node[i&(NODE_END_IDX)]
	} else {
		return notFound
	}
}

func (t *TransientVector) AssocN(i int, val interface{}) *TransientVector {
	t.ensureEditable()
	if i >= 0 && i < t.cnt {
		if i >= t.tailoff() {
			t.tail[i&(NODE_END_IDX)] = val
			return t
		}
		t.root = t.doAssoc(t.shift, t.root, i, val)
		return t
	}
	if i == t.cnt {
		return t.Conj(val)
	}
	panic(IndexOutOfBoundsException)
}

// Associate a new value at the given key. Key must be an integer.
func (t *TransientVector) Assoc(key interface{}, val interface{}) *TransientVector {
	switch k := key.(type) {
	case int:
		return t.AssocN(k, val)
	}
	panic("Key must be integer")
}

func (t *TransientVector) doAssoc(level uint, node *Node, i int, val interface{}) *Node {
	node = t.ensureEditableNode(node)
	ret := node
	if level == 0 {
		ret.array[i&(NODE_END_IDX)] = val
	} else {
		subidx := (i >> level) & (NODE_END_IDX)
		ret.array[subidx] = t.doAssoc(level-VECTOR_SHIFT, node.array[subidx].(*Node), i, val)
	}
	return ret
}

func (t *TransientVector) Pop() *TransientVector {
	t.ensureEditable()
	if t.cnt == 0 {
		panic("Can't pop empty vector")
	}
	if t.cnt == 1 {
		t.cnt = 0
		return t
	}
	i := t.cnt - 1
	if (i & (NODE_END_IDX)) > 0 {
		t.cnt--
		return t
	}

	newtail := t.editableArrayFor(t.cnt - 2)
	newroot := t.popTail(t.shift, t.root)
	newshift := t.shift
	if newroot == nil {
		newroot = &Node{edit: t.root.edit}
	}
	if t.shift > VECTOR_SHIFT && newroot.array[1] == nil {
		newroot = t.ensureEditableNode(newroot.array[0].(*Node))
		newshift -= VECTOR_SHIFT
	}
	t.root = newroot
	t.shift = newshift
	t.cnt--
	t.tail = newtail
	return t
}

func (t *TransientVector) popTail(level uint, node *Node) *Node {
	node = t.ensureEditableNode(node)
	subidx := ((t.cnt - 2) >> level) & (NODE_END_IDX)
	if level > VECTOR_SHIFT {
		newchild := t.popTail(level-VECTOR_SHIFT, node.array[subidx].(*Node))
		if &newchild == nil && subidx == 0 {
			return nil
		} else {
			ret := node
			ret.array[subidx] = newchild
			return ret
		}
	} else if subidx == 0 {
		return nil
	} else {
		ret := node
		ret.array[subidx] = nil
		return ret
	}
}
