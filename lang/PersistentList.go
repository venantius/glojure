package lang

import (
	"container/list" // linked list
)

/* Declaration block: Primordial */

type Primordial struct {
	*RestFn
}

func (p *Primordial) GetRequiredArity() int {
	return 0
}

func (p *Primordial) doInvoke(args interface{}) interface{} {
	switch a := args.(type) {
	case ArraySeq:
		argsarray := a.array
		var ret IPersistentList
		ret = EMPTY_PERSISTENT_LIST
		for i := len(argsarray) - 1; i >= 0; i-- {
			ret = ret.Cons(argsarray[i]).(IPersistentList)
		}
		return ret
	}
	list := list.New()
	for s := RT.Seq(args); s != nil; s = s.Next() {
		list.PushBack(s.First())
	}
	return create(list)
}

func (p *Primordial) InvokeStatic(args ISeq) interface{} {
	switch a := args.(type) {
	case *ArraySeq:
		argsarray := a.array
		var ret IPersistentList
		ret = EMPTY_PERSISTENT_LIST
		for i := len(argsarray) - 1; i >= 0; i-- {
			ret = ret.Cons(argsarray[i]).(IPersistentList)
		}
		return ret
	}
	list := list.New()
	for s := RT.Seq(args); s != nil; s = s.Next() {
		list.PushBack(s.First())
	}
	return create(list)
}

func (p *Primordial) WithMeta(meta IPersistentMap) IObj {
	panic(UnsupportedOperationException)
}

func (p *Primordial) Meta() IPersistentMap {
	return nil
}

/* Declaration block: PersistentList */

// NOTE: Implements IPersistentList, IReduce, List, Counted
type PersistentList struct {
	*ASeq

	_meta IPersistentMap // Inherited from Obj -> ASeq

	_first interface{}
	_rest  IPersistentList
	_count int
}

func create(init *list.List) IPersistentList {
	var ret IPersistentList
	ret = EMPTY_PERSISTENT_LIST
	for i := init.Back(); i != nil; i = i.Prev() {
		ret = ret.Cons(i).(IPersistentList)
	}
	return ret
}

func (l *PersistentList) First() interface{} {
	return l._first
}

func (l *PersistentList) Next() ISeq {
	if l._count == 1 {
		return nil
	}
	return l._rest.(ISeq)
}

func (l *PersistentList) Peek() interface{} {
	return l.First()
}

func (l *PersistentList) Pop() IPersistentStack {
	if l._rest == nil {
		return EMPTY_PERSISTENT_LIST.WithMeta(l._meta)
	}
	return l._rest
}

func (l *PersistentList) Count() int {
	return l._count
}

func (l *PersistentList) Cons(i interface{}) IPersistentCollection {
	return &PersistentList{
		_meta:  l.Meta(),
		_first: i,
		_rest:  l,
		_count: l._count + 1,
	}
}

func (l *PersistentList) Empty() IPersistentCollection {
	return EMPTY_PERSISTENT_LIST.WithMeta(l.Meta())
}

func (l *PersistentList) WithMeta(meta IPersistentMap) *PersistentList {
	if meta != l._meta {
		return &PersistentList{
			_meta:  meta,
			_first: l._first,
			_rest:  l._rest,
			_count: l._count,
		}
	}
	return l
}

func (l *PersistentList) ReduceWithInit(f IFn, start interface{}) interface{} {
	ret := f.Invoke(start, l.First())
	for s := l.Next(); s != nil; s = s.Next() {
		if RT.IsReduced(ret) {
			return ret.(IDeref).Deref()
		}
		ret = f.Invoke(ret, s.First())
	}
	if RT.IsReduced(ret) {
		return ret.(IDeref).Deref()
	}
	return ret
}

func (l *PersistentList) Reduce(f IFn) interface{} {
	ret := l.First()
	for s := l.Next(); s != nil; s = s.Next() {
		ret = f.Invoke(ret, s.First())
		if RT.IsReduced(ret) {
			return ret.(IDeref).Deref()
		}
	}
	return ret
}

/* Declaration Block: EmptyList */

// NOTE: Implements IPersistentList, List, ISeq, Counted, IHashEq
type EmptyList struct {
	*Obj
	PersistentList

	// inherit from Obj
	_meta IPersistentMap

	// inherit from PersistentList
	_first interface{}
	_rest  IPersistentList
	_count int

	hasheq int // Java default is Murmur3.hashOrdered(Collections.EMPTY_LIST)
}

var EMPTY_PERSISTENT_LIST = &EmptyList{}

func (e *EmptyList) HashCode() int {
	return 1
}

func (e *EmptyList) HashEq() int {
	return e.hasheq
}

func (e *EmptyList) String() string {
	return "()"
}

func (e *EmptyList) Equals(i interface{}) bool {
	var b bool
	switch i.(type) {
	case Sequential:
		b = true
	case List:
		b = true
	default:
		b = false
	}
	return b && RT.Seq(i) == nil
}

func (e *EmptyList) Equiv(i interface{}) bool {
	return e.Equals(i)
}

func (e *EmptyList) First() interface{} {
	return nil
}

func (e *EmptyList) Next() ISeq {
	return nil
}

func (e *EmptyList) More() ISeq {
	return e
}

func (e *EmptyList) Cons(i interface{}) IPersistentCollection {
	return &PersistentList{
		_meta:  e.Meta(),
		_first: i,
		_rest:  nil,
		_count: 1,
	}
}

func (e *EmptyList) Empty() IPersistentCollection {
	return e
}

func (e *EmptyList) WithMeta(meta IPersistentMap) *EmptyList {
	if meta != e.Meta() {
		return &EmptyList{_meta: meta}
	}
	return e
}

func (e *EmptyList) Peek() interface{} {
	return nil
}

func (e *EmptyList) Pop() IPersistentStack {
	panic("Can't pop empty list")
}

func (e *EmptyList) Count() int {
	return 0
}

func (e *EmptyList) Seq() ISeq {
	return nil
}

func (e *EmptyList) Size() int {
	return 0
}

func (e *EmptyList) IsEmpty() bool {
	return true
}

func (e *EmptyList) Contains(i interface{}) bool {
	return false
}

// TODO: Iterator code goes here

// NOTE: Overloaded
func (e *EmptyList) ToArray(args ...[]interface{}) []interface{} {
	if len(args) == 0 {
		return RT.EMPTY_ARRAY()
	} else {
		if len(args[0]) > 0 {
			args[0][0] = nil
		}
		return args[0]
	}
}

func (e *EmptyList) Add(i interface{}) bool {
	panic("Unsupported Operation")
}

func (e *EmptyList) Remove(i interface{}) bool {
	panic("Unsupported Operation")
}

func (e *EmptyList) AddAll(collection Collection) bool {
	panic("Unsupported Operation")
}

func (e *EmptyList) Clear() {
	panic("Unsupported Operation")
}

func (e *EmptyList) RetainAll(collection Collection) bool {
	panic("Unsupported Operation")
}

func (e *EmptyList) RemoveAll(collection Collection) bool {
	panic("Unsupported Operation")
}

func (e *EmptyList) ContainsAll(collection Collection) bool {
	return collection.IsEmpty()
}
