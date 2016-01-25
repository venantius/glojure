package lang
import (
	"fmt"
	"reflect"
)

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

	Equals(o interface{}) bool
	HashCode() int
	IndexOf(o interface{}) int
	// Iterator() Iterator // TODO: don't want to deal with this
	LastIndexOf(o interface{}) int
	// ListIterator(i int) List // TODO: don't want to deal with this right now.
	String() string
}

func APersistentVector_String(a APersistentVector) string {
	// return RT.PrintString(a)
	return RT.PrintString(a)
}

func APersistentVector_Seq(a APersistentVector) ISeq {
	if a.Count() > 0 {
		return &APersistentVectorSeq{
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
		fmt.Println("Here!")
		fmt.Println(v, i)
		for i := 0; i < v.Count(); i++ {
			if !Util.Equals(v.Nth(i, nil), it.Nth(i, nil)) {
				fmt.Println(v.Nth(i, nil), it.Nth(i, nil))
				fmt.Println(reflect.TypeOf(v.Nth(i, nil)), reflect.TypeOf(it.Nth(i, nil)))
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

func APersistentVector_Length(a APersistentVector) int {
	return a.Count()
}

func APersistentVector_Nth(a APersistentVector, i int, notFound interface{}) interface{} {
	if i >= 0 && i < a.Count() {
		return a.Nth(i, nil)
	}
	return notFound
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

func APersistentVector_EntryAt(a APersistentVector, key interface{}) IMapEntry {
	if IsInt(key) {
		var i int = key.(int)
		if i >= 0 && i < a.Count() {
			return CreateMapEntry(key, a.Nth(i, nil))
		}
	}
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

type APersistentVectorSeq struct {
	*ASeq

	_meta IPersistentMap
	v     IPersistentVector
	i     int
}

func (s *APersistentVectorSeq) First() interface{} {
	return s.v.Nth(s.i, nil)
}

func (s *APersistentVectorSeq) Next() ISeq {
	if (s.i + 1) < s.v.Count() {
		return &APersistentVectorSeq{
			v: s.v,
			i: s.i + 1,
		}
	}
	return nil
}

func (s *APersistentVectorSeq) Index() int {
	return s.i
}

func (s *APersistentVectorSeq) Count() int {
	return s.v.Count() - s.i
}

func (s *APersistentVectorSeq) WithMeta(meta IPersistentMap) *APersistentVectorSeq {
	return &APersistentVectorSeq{
		_meta: meta,
		v:     s.v,
		i:     s.i,
	}
}

// TODO
func (s *APersistentVectorSeq) Reduce(f IFn, start interface{}) interface{} {
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

func (r *SubVector) Assoc(key interface{}, val interface{}) Associative {
	return APersistentVector_Assoc(r, key, val)
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

func (r *SubVector) HashCode() int {
	return APersistentVector_HashCode(r)
}

func (r *SubVector) HashEq() int {
	return APersistentVector_HashEq(r)
}

func (r *SubVector) IndexOf(o interface{}) int {
	return APersistentVector_IndexOf(r, o)
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

func (r *SubVector) Seq() ISeq {
	return APersistentVector_Seq(r)
}

func (r *SubVector) String() string {
	return APersistentVector_String(r)
}

func (r *SubVector) ValAt(key interface{}, notFound interface{}) interface{} {
	return APersistentVector_ValAt(r, key, notFound)
}