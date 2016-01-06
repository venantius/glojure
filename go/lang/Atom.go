package lang

import (
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

// NOTE: Implements IAtom
type Atom struct {
	*ARef

	_meta IPersistentMap
	state interface{} // AtomicReference
}

func (a *Atom) Deref() interface{} {
	x := a.state.(unsafe.Pointer)
	return atomic.LoadPointer(&x)
}

// TODO
func (a *Atom) Swap(f IFn, args ...interface{}) interface{} {
	return nil
}

func (a *Atom) CompareAndSet(oldv interface{}, newv interface{}) bool {
	a.validate(nil, newv)
	state := a.state.(unsafe.Pointer)
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
	p := a.state.(unsafe.Pointer)
	atomic.StorePointer(&p, newval.(unsafe.Pointer))
	a.NotifyWatches(oldval, newval)
	return newval
}
