package lang

// TODO: Extends AFn
// NOTE: Implements IPersistentVector, Iterable, List, RandomAccess, COmparable, Serializable, IHashEq
type APersistentVector struct {
	*AFn

	_hash   int
	_hasheq int
}

// TODO
func (a *APersistentVector) String() string {
	// return RT.PrintString(a)
	return ""
}

func (a *APersistentVector) Seq() ISeq {
	if a.Count() > 0 {
		return PersistentVectorSeq{} // TODO: Fix initializer
	}
	return nil
}

func (a *APersistentVector) RSeq() ISeq {
	if a.Count() > 0 {
		return PersistentVectorRSeq{} // TODO: fix initializer
	}
	return nil
}

// TODO
func doEquals(v IPersistentVector, i interface{}) bool {
	return true
}

// TODO
func doEquiv(v IPersistentVector, i interface{}) bool {
	return true
}

func (a *APersistentVector) Equals(i interface{}) bool {
	if i == a {
		return true
	}
	return doEquals(a, i)
}

func (a *APersistentVector) Equiv(i interface{}) bool {
	if i == a {
		return true
	}
	return doEquiv(a, i)
}

// TODO
func (a *APersistentVector) HashCode() int {
	return 0
}

// TODO
func (a *APersistentVector) HashEq() int {
	return 0
}

func (a *APersistentVector) Get(index int) interface{} {
	return a.Nth(index)
}

func (a *APersistentVector) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < a.Count() {
		return a.Nth(i)
	}
	return notFound
}

func (a *APersistentVector) Remove(i int) interface{} {
	// TODO
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) IndexOf(interface{}) int {
	for i := 0; i < a.Count(); i++ {
		// if Util.equiv...TODO
		i++
	}
	return -1
}

// TODO
func (a *APersistentVector) LastIndexOf(i interface{}) int {
	return nil
}

// TODO
// func (a *APersistentVector) ListIterator() TYPE {
// }

// TODO
// func (a *APersistentVector) RangedIterator() TYPE {
// }

func (a *APersistentVector) SubList(fromIndex int, toIndex int) List {
	return RT.Subvec(a, fromIndex, toIndex).(List)
}

func (a *APersistentVector) Set(i int, o interface{}) interface{} {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) Add(i int, o interface{}) interface{} {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) AddAll(i int, c Collection) bool {
	panic(UnsupportedOperationException)
}

// TODO
func (a *APersistentVector) Invoke(arg1 interface{}) interface{} {
	return nil
}

// TODO
func (a *APersistentVector) Iterator() Iterator {
	return nil
}

func (a *APersistentVector) Peek() interface{} {
	if a.Count() > 0 {
		return a.Nth(a.Count() - 1)
	}
	return nil
}

// TODO
func (a *APersistentVector) ContainsKey(key interface{}) bool {
	return true
}

// TODO
func (a *APersistentVector) EntryAt(key interface{}) IMapEntry {
	return nil
}

// TODO
func (a *APersistentVector) Assoc(key interface{}, val interface{}) IPersistentVector {
	return nil
}

// TODO
func (a *APersistentVector) ValAt(key interface{}, notFound interface{}) interface{} {
	return nil
}

// NOTE: There's a note in here about everything else here being "java.util.Collection implementation"

func (a *APersistentVector) ToArray() []interface{} {
	ret := make([]interface{}, a.Count())
	for i := 0; i < a.Count(); i++ {
		ret[i] = a.Nth(i)
	}
	return ret
}

// NOTE: Commented out this implementation
// func (a *APersistentVector) Add(o interface{}) bool {
// 	panic(UnsupportedOperationException)
// }

type PersistentVectorSeq struct{} // TODO

type PersistentVectorRSeq struct{} // TODO
