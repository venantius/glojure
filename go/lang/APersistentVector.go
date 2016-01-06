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
		return &PersistentVectorSeq{} // TODO: Fix initializer
	}
	return nil
}

func (a *APersistentVector) RSeq() ISeq {
	if a.Count() > 0 {
		return &APersistentVectorRSeq{} // TODO: fix initializer
	}
	return nil
}

func (a *APersistentVector) Count() int {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) Cons(i interface{}) IPersistentCollection {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) Empty() IPersistentCollection {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) Length() int {
	panic(UnsupportedOperationException)
}

func (a *APersistentVector) Pop() IPersistentStack {
	panic(UnsupportedOperationException)
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
	return a.Nth(index, nil)
}

func (a *APersistentVector) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < a.Count() {
		return a.Nth(i, nil)
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
	return 0
}

// TODO
// func (a *APersistentVector) ListIterator() TYPE {
// }

// TODO
// func (a *APersistentVector) RangedIterator() TYPE {
// }

/* NOTE: uncomment
func (a *APersistentVector) SubList(fromIndex int, toIndex int) List {
	return RT.SubVec(a, fromIndex, toIndex).(List)
}
*/

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
func (a *APersistentVector) Iterator() *Iterator {
	return nil
}

func (a *APersistentVector) Peek() interface{} {
	if a.Count() > 0 {
		return a.Nth(a.Count()-1, nil)
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
func (a *APersistentVector) Assoc(key interface{}, val interface{}) Associative {
	return nil
}

func (a *APersistentVector) AssocN(i int, val interface{}) IPersistentVector {
	panic(UnsupportedOperationException)
}

// TODO
func (a *APersistentVector) ValAt(key interface{}, notFound interface{}) interface{} {
	return nil
}

// NOTE: There's a note in here about everything else here being "java.util.Collection implementation"

func (a *APersistentVector) ToArray() []interface{} {
	ret := make([]interface{}, a.Count())
	for i := 0; i < a.Count(); i++ {
		ret[i] = a.Nth(i, nil)
	}
	return ret
}

// NOTE: Commented out this implementation
// func (a *APersistentVector) Add(o interface{}) bool {
// 	panic(UnsupportedOperationException)
// }

// Declaration block: PersistentVectorSeq

type PersistentVectorSeq struct {
	*ASeq

	_meta IPersistentMap
	v     IPersistentVector
	i     int
}

func (s *PersistentVectorSeq) First() interface{} {
	return s.v.Nth(s.i, nil)
}

func (s *PersistentVectorSeq) Next() ISeq {
	if (s.i + 1) < s.v.Count() {
		return &PersistentVectorSeq{
			v: s.v,
			i: s.i + 1,
		}
	}
	return nil
}

func (s *PersistentVectorSeq) Index() int {
	return s.i
}

func (s *PersistentVectorSeq) Count() int {
	return s.v.Count() - s.i
}

func (s *PersistentVectorSeq) WithMeta(meta IPersistentMap) *PersistentVectorSeq {
	return &PersistentVectorSeq{
		_meta: meta,
		v:     s.v,
		i:     s.i,
	}
}

// TODO
func (s *PersistentVectorSeq) Reduce(f IFn, start interface{}) interface{} {
	return nil
}

// NOTE: Implements IndexedSeq, Counted
type APersistentVectorRSeq struct {
	*ASeq

	_meta IPersistentMap
	v     IPersistentVector
	i     int
}

func (r *APersistentVectorRSeq) First() interface{} {
	return r.v.Nth(r.i, nil)
}

func (r *APersistentVectorRSeq) Next() ISeq {
	if r.i > 0 {
		return &APersistentVectorRSeq{
			v: r.v,
			i: r.i - 1,
		}
	}
	return nil
}

func (r *APersistentVectorRSeq) Index() int {
	return r.i
}

func (r *APersistentVectorRSeq) Count() int {
	return r.i + 1
}

func (r *APersistentVectorRSeq) WithMeta(meta IPersistentMap) *APersistentVectorRSeq {
	return &APersistentVectorRSeq{
		_meta: meta,
		v:     r.v,
		i:     r.i,
	}
}

// NOTE: Implements IObj
type SubVector struct {
	*APersistentVector

	_meta IPersistentMap
	v     IPersistentVector
	start int
	end   int
}

// TODO: Custom SubVector initializer?

// TODO
func (r *SubVector) Iterator() *Iterator {
	return nil
}

func (r *SubVector) Nth(i int, notFound interface{}) interface{} {
	if (r.start+i) >= r.end || i < 0 {
		panic(IndexOutOfBoundsException)
	}
	return r.v.Nth(r.start+i, notFound)
}

func (r *SubVector) AssocN(i int, val interface{}) IPersistentVector {
	if (r.start + i) > r.end {
		panic(IndexOutOfBoundsException)
	} else if r.start+i == r.end {
		return r.Cons(val).(IPersistentVector)
	}
	return &SubVector{
		_meta: r._meta,
		v:     r.v.AssocN(r.start+i, val),
		start: r.start,
		end:   r.end,
	}
}

func (r *SubVector) Count() int {
	return r.end - r.start
}

func (r *SubVector) Cons(o interface{}) IPersistentCollection {
	return &SubVector{
		_meta: r._meta,
		v:     r.v.AssocN(r.end, o),
		start: r.start,
		end:   r.end,
	}
}

func (r *SubVector) Empty() IPersistentCollection {
	return nil
	// return EMPTY_PERSISTENT_VECTOR.WithMeta(r.Meta()).(IPersistentCollection) // TODO: This is probably cheating.
}

func (r *SubVector) Pop() IPersistentStack {
	if (r.end - 1) == r.start {
		return nil
		// return EMPTY_PERSISTENT_VECTOR.(IPersistentCollection) // TODO cheating again
	}
	return &SubVector{
		_meta: r._meta,
		v:     r.v,
		start: r.start,
		end:   r.end - 1,
	}
}

func (r *SubVector) WithMeta(meta IPersistentMap) *SubVector {
	if meta == r._meta {
		return r
	}
	return &SubVector{
		_meta: meta,
		v:     r.v,
		start: r.start,
		end:   r.end,
	}
}

func (r *SubVector) Meta() IPersistentMap {
	return r._meta
}
