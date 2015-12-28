package lang

import (
	"errors"
	"fmt"     // really just for debugging
	"reflect" // this too.
)

var indexOutOfBoundsException = errors.New("Index out of bounds.")
var emptyVectorPopError = errors.New("Can't pop empty vector.")

// TODO: Remove this
type Iterator struct{}
type Iterable interface{}
type List struct{}

// TODO:
// PersistentVector extends APersistentVector, implements IObj, IEditableCollection, IKVReduce
type PersistentVector struct {
	*APersistentVector

	cnt   int // count
	shift uint
	root  Node
	tail  []interface{}
	_meta IPersistentMap
}

// TODO: Implements Serializable
type Node struct {
	// TODO: edit AtomicReference ??
	edit  interface{}
	array []interface{}
}

const (
	VECTOR_SHIFT = 5
	NODE_SIZE    = 32
)

var EMPTY_NODE = Node{array: make([]interface{}, NODE_SIZE)}

var EMPTY = PersistentVector{cnt: 0, shift: VECTOR_SHIFT, root: EMPTY_NODE, tail: make([]interface{}, 0)}

// TODO: IFn TRANSIENT_VECTOR_CONJ

// Return a new PersistentVector with the items passed in.
func adopt(items []interface{}) PersistentVector {
	return PersistentVector{
		cnt:   len(items),
		shift: VECTOR_SHIFT,
		root:  EMPTY_NODE,
		tail:  items,
	}
}

// TODO
func createVectorFromIReduceInit(items IReduceInit) PersistentVector {
	return PersistentVector{}
}

// TODO
func createVectorFromISeq(items ISeq) PersistentVector {
	return PersistentVector{}
}

// TODO
func createVectorFromList(list List) PersistentVector {
	return PersistentVector{}
}

// TODO
func createVectorFromIterable(items Iterable) PersistentVector {
	return PersistentVector{}
}

// TODO
func createVectorFromInterfaceSlice(items ...interface{}) PersistentVector {
	return PersistentVector{}
}

// TODO: UNFINISHED
func CreateVector(items ...interface{}) PersistentVector {
	ret := *&EMPTY
	fmt.Println(reflect.TypeOf(items[0]))
	switch items[0].(type) {
	case IReduceInit:
		createVectorFromIReduceInit(items[0].(IReduceInit))
	case ISeq:
		createVectorFromISeq(items[0].(ISeq))
	case Iterable:
		createVectorFromIterable(items[0].(Iterable))
	default:
		createVectorFromInterfaceSlice(items)
	}
	return ret
}

// TODO: public TransientVector asTransient()

// TODO: Check this.
func (v *PersistentVector) tailoff() int {
	if v.cnt < NODE_SIZE {
		return 0
	}
	// NOTE: I'm not totally clear on whether this is consistent with the Java version.
	return ((v.cnt - 1) >> VECTOR_SHIFT) << VECTOR_SHIFT
}

// TODO: Check this.
func (v *PersistentVector) ArrayFor(i int) []interface{} {
	if i < 0 || i >= v.cnt {
		panic(indexOutOfBoundsException)
	}
	if i >= v.tailoff() {
		return v.tail
	}
	n := *&v.root
	for level := v.shift; level > 0; level -= VECTOR_SHIFT {
		// NOTE: bitshift is probably wrong here as well.
		n = n.array[(i>>level)&(NODE_SIZE-1)].(Node)
	}
	return n.array
}

// Return the number of items in the vector.
func (v *PersistentVector) Count() int {
	return v.cnt
}

func (v *PersistentVector) nth(i int) interface{} {
	subsl := v.ArrayFor(i)
	return subsl[i&(NODE_SIZE-1)]
}

// Retrieve the nth item in the vector. If the index being retrieved is beyond
// the length of the vector, returns the notFound value.
func (v *PersistentVector) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < v.cnt {
		return v.nth(i)
	}
	return notFound
}

// TODO: Follow up on this
// NOTE: Is this even needed?
func (v *PersistentVector) String() string {
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

func (v *PersistentVector) AssocN(i int, val interface{}) PersistentVector {
	if i >= 0 && i < v.cnt {
		if i >= v.tailoff() {
			newTail := make([]interface{}, len(v.tail))
			copy(newTail, v.tail)
			newTail[i&(NODE_SIZE-1)] = val
			return PersistentVector{
				_meta: v.Meta(),
				cnt:   v.cnt,
				shift: v.shift,
				root:  v.root,
				tail:  newTail,
			}
		}
		return PersistentVector{
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
	panic(indexOutOfBoundsException)
}

func doAssoc(level uint, node Node, i int, val interface{}) Node {
	var arr []interface{}
	copy(arr, node.array)
	ret := Node{edit: node.edit, array: arr}
	if level == 0 {
		ret.array[i&(NODE_SIZE-1)] = val
	} else {
		// NOTE: Bitwise issues again.
		subidx := (i >> level) & (NODE_SIZE - 1)
		ret.array[subidx] = doAssoc(level-VECTOR_SHIFT, node.array[subidx].(Node), i, val)
	}
	return ret
}

// Return a new PersistentVector with new metadata.
func (v *PersistentVector) WithMeta(meta IPersistentMap) PersistentVector {
	return PersistentVector{
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

func (v *PersistentVector) newPath(edit interface{}, level uint, node Node) Node {
	if level == 0 {
		return node
	}
	ret := Node{edit: edit}
	ret.array[0] = v.newPath(edit, level-VECTOR_SHIFT, node)
	return ret
}

func (v *PersistentVector) pushTail(level uint, parent Node, tailnode Node) Node {
	// NOTE: bitshifts require review
	subidx := ((v.cnt - 1) >> level) & NODE_SIZE
	ret := Node{edit: parent.edit, array: *&parent.array}
	nodeToInsert := Node{}
	if level == VECTOR_SHIFT {
		nodeToInsert = tailnode
	} else {
		child := parent.array[subidx]
		if child != nil {
			nodeToInsert = v.pushTail(level-VECTOR_SHIFT, child.(Node), tailnode)
		} else {
			nodeToInsert = v.newPath(v.root.edit, level-VECTOR_SHIFT, tailnode)
		}
	}
	ret.array[subidx] = nodeToInsert
	return ret
}

func (v *PersistentVector) Cons(val interface{}) PersistentVector {
	if v.cnt-v.tailoff() < NODE_SIZE {
		newTail := make([]interface{}, len(v.tail)+1)
		copy(newTail, v.tail)
		newTail[len(v.tail)] = val
		return PersistentVector{
			_meta: v.Meta(),
			cnt:   v.cnt + 1,
			shift: v.shift,
			root:  v.root,
			tail:  newTail,
		}
	}
	newroot := Node{}
	tailnode := Node{v.root.edit, v.tail}
	newshift := v.shift
	fmt.Println(tailnode)
	// NOTE: Again, not comfortable with bit shifting here.
	if (v.cnt >> VECTOR_SHIFT) > (1 << v.shift) {
		newroot = Node{edit: v.root.edit} // defaults?
		newroot.array[0] = v.root
		newroot.array[1] = v.newPath(v.root.edit, v.shift, tailnode)
		newshift += 5
	} else {
		newroot = v.pushTail(v.shift, v.root, tailnode)
	}
	return PersistentVector{
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
	// TODO: Verify that this actually works.
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
			if IsReduced(init) {
				return init.(IDeref).Deref()
			}
		}
		step = len(array)
	}
	return init
}

// TODO
func (v *PersistentVector) KVReduce(f IFn, init interface{}) interface{} {
	return nil
}

// NOTE: implements IChunkedSeq, Counted
type ChunkedSeq struct {
	*ASeq

	vec    PersistentVector
	node   []interface{}
	i      int
	offset int
}

// TODO
func (c *ChunkedSeq) ChunkedFirst() IChunk {
	return nil
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
		return &EMPTY_PERSISTENT_LIST
	}
	return s
}

func (c *ChunkedSeq) WithMeta(meta IPersistentMap) ChunkedSeq {
	if meta == c.vec._meta {
		return *c
	}
	return ChunkedSeq{
		vec:    c.vec.WithMeta(meta),
		node:   c.node,
		i:      c.i,
		offset: c.offset,
	}
}

func (c *ChunkedSeq) First() interface{} {
	return c.node[c.offset]
}

// TODO: This won't work until we implement Cons
func (c *ChunkedSeq) Next() ISeq {
	if c.offset+1 < len(c.node) {
		return nil
		/*
			ChunkedSeq{
				vec:    c.vec,
				node:   c.node,
				i:      c.i,
				offset: c.offset + 1,
			}
		*/
	}
	return c.ChunkedNext()
}

// Return the size of the chunked sequence.
func (c *ChunkedSeq) Count() int {
	return c.vec.cnt - (c.i + c.offset)
}

// Empty the vector's contents.
func (v *PersistentVector) Empty() PersistentVector {
	return EMPTY.WithMeta(v.Meta())
}

// TODO
func (v *PersistentVector) Pop() PersistentVector {
	if v.cnt == 0 {
		panic(emptyVectorPopError)
	}
	if v.cnt == 1 {
		return EMPTY.WithMeta(v.Meta())
	}
	if (v.cnt - v.tailoff()) > 1 {
		newTail := make([]interface{}, len(v.tail)-1)
		copy(v.tail, newTail) // NOTE: Figure this out
		return PersistentVector{
			_meta: v.Meta(),
			cnt:   v.cnt - 1,
			shift: v.shift,
			root:  v.root,
			tail:  newTail,
		}
	}
	newTail := v.ArrayFor(v.cnt - 2)

	_newroot := v.popTail(v.shift, v.root)
	newroot := *_newroot

	newshift := v.shift
	if &_newroot == nil {
		newroot = EMPTY_NODE
	}
	if v.shift > VECTOR_SHIFT && &newroot.array[1] == nil {
		newroot = newroot.array[0].(Node)
		newshift -= VECTOR_SHIFT
	}
	return PersistentVector{
		_meta: v.Meta(),
		cnt:   v.cnt - 1,
		shift: newshift,
		root:  newroot,
		tail:  newTail,
	}
}

func (v *PersistentVector) popTail(level uint, node Node) *Node {
	subidx := ((v.cnt - 2) >> level) & (NODE_SIZE - 1)
	if level > VECTOR_SHIFT {
		newchild := v.popTail(level-VECTOR_SHIFT, node.array[subidx].(Node))
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

// TODO: TransientVector
type TransientVector struct {
	cnt   int
	shift int
	root  Node
	tail  []interface{}
}

// TODO...the rest of this
