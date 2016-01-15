package lang

import (
	"sync"
	"bytes"
)

/*
	Assorted class variables
 */

// TODO: I found this really useful: http://www.golangbootcamp.com/book/concurrency

var privateKey *Keyword = InternKeywordByNsName("private")
var privateMeta IPersistentMap = &PersistentArrayMap{
	array: []interface{}{privateKey, true},
}
var macroKey = InternKeywordByNsName("macro")
var nameKey = InternKeywordByNsName("name")
var nsKey = InternKeywordByNsName("ns")

/*
	Var

	Implements: IFn, IRef, Settable
 */

type Var struct {
	ARef

	rev int
	privateMeta IPersistentMap
}

/*
	VarTBox [TBox]
 */

type VarTBox struct {
	val interface{}
	lock *sync.Mutex // instead of keeping track of thread
}

/*
	VarUnbound [Unbound]
 */

type VarUnbound struct {
	AFn

	v *Var
}

func (u *VarUnbound) String() string {
	var b bytes.Buffer
	b.Write("Unbound: ")
	b.Write(u.v.String())
	return b.String()
}

func (u *VarUnbound) ThrowArity(n int) {
	var b bytes.Buffer
	b.Write("Attempting to call unbound fn: ")
	b.Write(u.v.String())
	panic(b.String())
}

/*
	VarFrame [Frame]
 */

type Frame struct {
	bindings Associative
	prev *Frame
}

var TOP_FRAME = &Frame{bindings: EMPTY_PERSISTENT_HASH_MAP, nil}

func (f *Frame) Clone() interface{} {
	return &Frame{
		bindings: f.bindings,
		prev: nil,
	}
}

/*
	VarFrameDvals [dvals](anonymous class)

	This anonymous class is just a ThreadLocal<Frame> in JVM Clojure with an initialiValue() func.
 */

type VarFrameDvals struct {
	Frame

	lock *sync.Mutex
}

func (d *VarFrameDvals) InitialValue() *Frame {
	return TOP_FRAME
}


