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
	// TODO: _meta IPersistentMap
}

// TODO: Implements Serializable
type Node struct {
	// TODO: edit AtomicReference ??
	array []interface{}
}

const (
	VECTOR_SHIFT = 5
	NODE_LENGTH  = 32
)

var EMPTY_NODE = Node{array: make([]interface{}, NODE_LENGTH)}

var EMPTY = PersistentVector{cnt: 0, shift: VECTOR_SHIFT, root: EMPTY_NODE, tail: make([]interface{}, 0)}

// TODO: IFn TRANSIENT_VECTOR_CONJ

func adopt(items []interface{}) {
	// TODO
}

// TODO: Follow up on this.
func Create(items ...interface{}) PersistentVector {
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
	if v.cnt < NODE_LENGTH {
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
	n := v.root
	for level := v.shift; level > 0; level -= VECTOR_SHIFT {
		n = n.array[(i>>level)&((VECTOR_SHIFT)-1)].(Node)
	}
	return n.array
}

// Return the number of items in the vector.
func (v *PersistentVector) Count() int {
	return v.cnt
}

func (v *PersistentVector) nth(i int) interface{} {
	subsl := v.arrayFor(i)
	return subsl[i&(NODE_LENGTH-1)]
}

func (v *PersistentVector) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < v.cnt {
		return v.nth(i)
	}
	subsl := v.arrayFor(i)
	return subsl[i&(NODE_LENGTH-1)]
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
