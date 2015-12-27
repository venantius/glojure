package lang

import (
	"errors"
	"fmt"     // really just for debugging
	"reflect" // this too.
)

var indexOutOfBoundsException = errors.New("Index out of bounds.")

// TODO:
// PersistentVector extends APersistentVector, implements IObj, IEditableCollection, IReduce, IKVReduce
type PersistentVector struct {
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

func adopt(items []interface{}) {
	// TODO
}

// TODO: UNFINISHED
func CreateVector(items ...interface{}) PersistentVector {
	ret := *&EMPTY
	fmt.Println(reflect.TypeOf(items[0]))
	switch items[0].(type) {
	case IReduceInit:
		fmt.Println("IReduceInit")
	default:
		fmt.Println("unknown")
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
func (v *PersistentVector) arrayFor(i int) []interface{} {
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
	subsl := v.arrayFor(i)
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

// TODO
func (v *PersistentVector) AssocN(i int, val interface{}) {
}

// TODO
func (v *PersistentVector) doAssoc(level int, node Node, i int, val interface{}) {
}

// TODO
func (v *PersistentVector) WithMeta(meta IPersistentMap) PersistentVector {
	return EMPTY
}

func (v *PersistentVector) Meta() IPersistentMap {
	return v._meta
}

func (v *PersistentVector) newPath(edit interface{}, level uint, node Node) Node {
	if level == 0 {
		return node
	}
	ret := Node{edit: edit}
	ret.array[0] = v.newPath(edit, level-5, node)
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

// TODO
func (v *PersistentVector) ChunkedSeq() IChunkedSeq {
	return nil
}

func (v *PersistentVector) Seq() ISeq {
	return v.ChunkedSeq()
}
