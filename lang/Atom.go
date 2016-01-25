package lang

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

/*
	This is Glojure's implementation of Clojure atoms. At the moment it is
	aggressively integrated with the sync/atomic library.

	I'm not sure, ultimately, if this is the right design. An alternative design
	that is potentially more idiomatic to Go might be to use a constructor that
	spawns goroutines for writer and reader channels, and where functions that
	attempt to mutate the underlying value have to do so through these channels.

	At the moment I've opted to avoid that implementation because it isn't clear
	to me that it's less complicated, and I believe it may be an incorrect
	design pattern for what Atoms are supposed to represent.

	I reserve the right to change my mind in the future.

	~ @venantius
*/

/*
	NOTE: Atoms must be initialized with the CreateAtom constructor. If not,
	they will have their watcher field set to nil, which will cause a runtime
	panic.
*/

// NOTE: Implements IAtom
type Atom struct {
	ARef
	validator IFn
	watches   int // TODO: should be IPersistentMap

	_meta IPersistentMap
	state unsafe.Pointer // AtomicReference
}

var EMPTY_PERSISTENT_MAP IPersistentMap

func (a *Atom) Initialize() *Atom {
	return &Atom{
		validator: a.validator,
		watches:   1,
		_meta:     a._meta,
		state:     a.state,
	}
}

func (a *Atom) Deref() interface{} {
	return atomic.LoadPointer(&a.state)
}

// TODO
func (a *Atom) Swap(f IFn, args ...interface{}) interface{} {
	return nil
}

func (a *Atom) CompareAndSet(oldv interface{}, newv interface{}) bool {
	a.validate(nil, newv)
	state := unsafe.Pointer(&a.state)
	ret := atomic.CompareAndSwapPointer(
		&state,
		oldv.(unsafe.Pointer),
		newv.(unsafe.Pointer),
	)
	if ret {
		a.NotifyWatches(oldv, newv)
	}
	return ret
}

func (a *Atom) Reset(newval interface{}) interface{} {
	oldval := a.Deref()
	a.validate(nil, newval)
	fmt.Println(newval)
	p := unsafe.Pointer(&a.state)
	atomic.StorePointer(&p, unsafe.Pointer(&newval))
	fmt.Println(a)
	a.NotifyWatches(oldval, newval)

	fmt.Println(a)
	return newval
}
