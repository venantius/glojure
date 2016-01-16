package lang

/*
	This was the first file where I decided to blatantly ignore the collection interface stuff from JVM Clojure.

	As Go lacks a comparable public interface for generics / collections, there's little to be gained by stubbing
	and duplicating behavior that has a clear functional alternative.
 */

/*
	APersistentSet

	Implements: IPersistentSet, Collection, Set, Serializable, IHashEq
 */

type APersistentSet struct {
	AFn

	_meta   IPersistentMap
	_hash   int
	_hasheq int
	impl    IPersistentMap
}

func (a *APersistentSet) String() string {
	return RT.PrintString(a)
}

func (a *APersistentSet) Contains(key interface{}) bool {
	return a.impl.ContainsKey(key)
}

func (a *APersistentSet) Get(key interface{}) interface{} {
	return a.impl.ValAt(key, nil)
}

func (a *APersistentSet) Count() int {
	return a.impl.Count()
}

func (a *APersistentSet) Seq() ISeq {
	return RT.Keys(a.impl)
}

func (a *APersistentSet) Invoke(args ...interface{}) interface{} {
	return a.Get(args[0])
}

func (a *APersistentSet) Equals(obj interface{}) bool {
	return SetEquals(a, obj)
}

func SetEquals(s1 IPersistentSet, obj interface{}) bool {
	if s1 == obj {
		return true
	}
	switch obj.(type) {
	case IPersistentSet:
	//
	default:
		return false
	}

	m := obj.(IPersistentSet)
	if m.Count() != s1.Count() {
		return false
	}

	for as := m.Seq(); as != nil; as = as.Next() {
		if !s1.Contains(as.First()) {
			return false
		}
	}
	return true
}

func (a *APersistentSet) Equiv(obj interface{}) bool {
	switch obj.(type) {
	case IPersistentSet:
	//
	default:
		return false
	}

	m := obj.(IPersistentSet)
	if m.Count() != a.Count() {
		return false
	}

	for as := m.Seq(); as != nil; as = as.Next() {
		if !a.Contains(as.First()) {
			return false
		}
	}
	return true
}

// TODO
func (a *APersistentSet) HashCode() int {
	return a._hash
}

// TODO
func (a *APersistentSet) HashEq() int {
	return a._hasheq
}

func (a *APersistentSet) ToArray() []interface{} {
	return RT.SeqToArray(a.Seq())
}

func (a *APersistentSet) ToPassedArray(arr []interface{}) []interface{} {
	return RT.SeqToPassedArray(a.Seq(), arr)
}

func (a *APersistentSet) Add(o interface{}) bool {
	panic(UnsupportedOperationException)
}

func (a *APersistentSet) Remove(o interface{}) bool {
	panic(UnsupportedOperationException)
}


func (a *APersistentSet) AddAll(args ...interface{}) bool {
	panic(UnsupportedOperationException)
}


func (a *APersistentSet) Clear() {
	panic(UnsupportedOperationException)
}


func (a *APersistentSet) RetainAll(o interface{}) {
	panic(UnsupportedOperationException)
}


func (a *APersistentSet) IsEmpty() bool {
	return a.Count() == 0
}

/*
	Required IPersistentSet methods
 */

func (a *APersistentSet) Cons(o interface{}) IPersistentCollection {
	panic(AbstractClassMethodException)
}

func (a *APersistentSet) Disjoin(o interface{}) IPersistentSet{
	panic(AbstractClassMethodException)
}

func (a *APersistentSet) Empty() IPersistentCollection {
	panic(AbstractClassMethodException)
}