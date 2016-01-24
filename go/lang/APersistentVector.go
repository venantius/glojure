package lang

/*
	APersistentVector

	Extends: AFn

	Note: In JVM clojure, this is an abstract class. This doesn't map to Go's inheritance
	model, and so instead I've chosen to instead make APersistentVector an interface, with
	associated static methods instead.

	There was probably a world in which I could have forced the original approach here of
	embedding structs to work, but not easily. I believe this approach should be much faster.

	APersistentVector had the new fields of _hash (int) and _hasheq (int)
 */

type APersistentVector interface {
	IPersistentVector
	IHashEq
	IFn

	Add(i int, o interface{})
	AddAll(i int, c Collection) bool
	Clear()
	CompareTo(o interface{}) int
	Contains(o interface{}) bool
	ContainsAll(o interface{}) bool
	Equals(o interface{}) bool
	Get(i int) interface{}
	HashCode() int
	IndexOf(o interface{}) int
	IsEmpty() bool
	// Iterator() Iterator // TODO: don't want to deal with this
	LastIndexOf(o interface{}) int
	// ListIterator(i int) List // TODO: don't want to deal with this right now.
	// Remove(o interface{}) bool // TODO: punting on this as well
	Set(i int, o interface{}) interface{}
	String() string
}

func APersistentVector_String(a APersistentVector) string {
	// return RT.PrintString(a)
	return RT.PrintString(a)
}

func APersistentVector_Seq(a APersistentVector) ISeq {
	if a.Count() > 0 {
		return &PersistentVectorSeq{
			v: a,
			i: 0,
		}
	}
	return nil
}

func APersistentVector_RSeq(a APersistentVector) ISeq {
	if a.Count() > 0 {
		return &APersistentVectorRSeq{
			v: a,
			i: a.Count() - 1,
		}
	}
	return nil
}

func APersistentVector_IsEmpty(a APersistentVector) bool {
	return a.Count() == 0
}

func APersistentVector_Contains(a APersistentVector, o interface{}) bool {
	for s := a.Seq(); s != nil; s = s.Next() {
		if Util.Equiv(s.First(), o) {
			return true
		}
	}
	return false
}

func APersistentVector_Length(a APersistentVector) int {
	return a.Count()
}

// TODO
func APersistentVector_CompareTo(a APersistentVector, o interface{}) int {
	return 0
}

func APersistentVector_doEquals(v IPersistentVector, i interface{}) bool {
	switch it := i.(type) {
	case IPersistentVector:
		if it.Count() != v.Count() {
			return false
		}
		for i := 0; i < v.Count(); i++ {
			if !Util.Equals(v.Nth(i, nil), it.Nth(i, nil)) {
				return false
			}
		}
		return true
	}
	// TODO: More type switches here
	return true
}

// TODO
func APersistentVector_doEquiv(v IPersistentVector, i interface{}) bool {
	return true
}

func APersistentVector_Equals(a APersistentVector, i interface{}) bool {
	if i == a {
		return true
	}
	return APersistentVector_doEquals(a, i)
}

func APersistentVector_Equiv(a APersistentVector, i interface{}) bool {
	if i == a {
		return true
	}
	return APersistentVector_doEquiv(a, i)
}

// TODO
func APersistentVector_HashCode(a APersistentVector) int {
	return 0
}

// TODO
func APersistentVector_HashEq(a APersistentVector) int {
	return 0
}

func APersistentVector_Get(a APersistentVector, index int) interface{} {
	return a.Nth(index, nil)
}

func APersistentVector_Nth(a APersistentVector, i int, notFound interface{}) interface{} {
	if i >= 0 && i < a.Count() {
		return a.Nth(i, nil)
	}
	return notFound
}

func APersistentVector_Remove(a APersistentVector, i int) interface{} {
	// TODO
	panic(UnsupportedOperationException)
}

func APersistentVector_IndexOf(a APersistentVector, o interface{}) int {
	for i := 0; i < a.Count(); i++ {
		if Util.Equiv(a.Nth(i, nil), o) {
			return i
		}
	}
	return -1
}

// TODO
func APersistentVector_LastIndexOf(a APersistentVector, i interface{}) int {
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

func APersistentVector_Set(a APersistentVector, i int, o interface{}) interface{} {
	panic(UnsupportedOperationException)
}

func APersistentVector_Add(a APersistentVector, i int, o interface{}) interface{} {
	panic(UnsupportedOperationException)
}

func APersistentVector_AddAll(a APersistentVector, i int, c Collection) bool {
	panic(UnsupportedOperationException)
}

func APersistentVector_Clear(a APersistentVector) {
	panic(UnsupportedOperationException)
}

// TODO
func APersistentVector_Invoke(a APersistentVector, arg1 interface{}) interface{} {
	return nil
}

// TODO
func APersistentVector_Iterator(a APersistentVector) *Iterator {
	return nil
}

func APersistentVector_Peek(a APersistentVector) interface{} {
	if a.Count() > 0 {
		return a.Nth(a.Count() - 1, nil)
	}
	return nil
}

// TODO
func APersistentVector_ContainsKey(a APersistentVector, key interface{}) bool {
	return true
}

// TODO
func APersistentVector_ContainsAll(a APersistentVector, c interface{}) bool {
	return true
}

// TODO
func APersistentVector_EntryAt(a APersistentVector, key interface{}) IMapEntry {
	return nil
}

// TODO
func APersistentVector_Assoc(a APersistentVector, key interface{}, val interface{}) Associative {
	return nil
}

func APersistentVector_AssocN(a APersistentVector, i int, val interface{}) IPersistentVector {
	panic(UnsupportedOperationException)
}

// TODO
func APersistentVector_ValAt(a APersistentVector, key interface{}, notFound interface{}) interface{} {
	return nil
}

// NOTE: There's a note in here about everything else here being "java.util.Collection implementation"

func APersistentVector_ToArray(a APersistentVector) []interface{} {
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

/*
	SubVector

	Implements: IObj, APersistentVector
 */

type SubVector struct {
	AFn // for now. We'll see.

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
	if (r.start + i) >= r.end || i < 0 {
		panic(IndexOutOfBoundsException)
	}
	return r.v.Nth(r.start + i, notFound)
}

func (r *SubVector) AssocN(i int, val interface{}) IPersistentVector {
	if (r.start + i) > r.end {
		panic(IndexOutOfBoundsException)
	} else if r.start + i == r.end {
		return r.Cons(val).(IPersistentVector)
	}
	return &SubVector{
		_meta: r._meta,
		v:     r.v.AssocN(r.start + i, val),
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

/*
	SubVector inheritance block: APersistentVector
 */

func (r *SubVector) Add(i int, o interface{}) {
	APersistentVector_Add(r, i, o)
}

func (r *SubVector) AddAll(i int, c Collection) bool {
	return APersistentVector_AddAll(r, i, c)
}

func (r *SubVector) Assoc(key interface{}, val interface{}) Associative {
	return APersistentVector_Assoc(r, key, val)
}

func (r *SubVector) Clear() {
	APersistentVector_Clear(r)
}

func (r *SubVector) CompareTo(o interface{}) int {
	return APersistentVector_CompareTo(r, o) // maybe?
}

func (r *SubVector) Contains(key interface{}) bool {
	return APersistentVector_Contains(r, key)
}

func (r *SubVector) ContainsAll(c interface{}) bool {
	return APersistentVector_ContainsAll(r, c)
}

func (r *SubVector) ContainsKey(key interface{}) bool {
	return APersistentVector_ContainsKey(r, key)
}

func (r *SubVector) EntryAt(key interface{}) IMapEntry {
	return APersistentVector_EntryAt(r, key)
}


func (r *SubVector) Equals(i interface{}) bool {
	return APersistentVector_Equals(r, i)
}

func (r *SubVector) Equiv(i interface{}) bool {
	return APersistentVector_Equiv(r, i)
}

func (r *SubVector) Get(index int) interface{} {
	return APersistentVector_Get(r, index)
}

func (r *SubVector) HashCode() int {
	return APersistentVector_HashCode(r)
}

func (r *SubVector) HashEq() int {
	return APersistentVector_HashEq(r)
}

func (r *SubVector) IndexOf(o interface{}) int {
	return APersistentVector_IndexOf(r, o)
}

func (r *SubVector) IsEmpty() bool {
	return APersistentVector_IsEmpty(r)
}

func (r *SubVector) LastIndexOf(i interface{}) int {
	return APersistentVector_LastIndexOf(r, i)
}

func (r *SubVector) Length() int {
	return APersistentVector_Length(r)
}

func (r *SubVector) Peek() interface{} {
	return APersistentVector_Peek(r)
}

func (r *SubVector) RSeq() ISeq {
	return APersistentVector_RSeq(r)
}

func (r *SubVector) Set(i int, o interface{}) interface {} {
	return APersistentVector_Set(r, i, o)
}

func (r *SubVector) Seq() ISeq {
	return APersistentVector_Seq(r)
}

func (r *SubVector) String() string {
	return APersistentVector_String(r)
}

func (r *SubVector) ValAt(key interface{}, notFound interface{}) interface{} {
	return APersistentVector_ValAt(r, key, notFound)
}