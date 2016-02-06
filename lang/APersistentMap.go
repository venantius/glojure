package lang

/*
	APersistentMap

	Extends: AFn

	Note: In JVM clojure, this is an abstract class. See APersistentVector for more details.

	APersistentMap had the new fields of _hash (int) and _hasheq (int)
 */

type APersistentMap interface {
	IPersistentMap
	MapEquivalence
	IHashEq
	IFn
	IEquals
}

func APersistentMap_Cons(a APersistentMap, obj interface{}) IPersistentCollection {
	switch o := obj.(type) {
	case IMapEntry: // NOTE: Map.Entry in Java
		return a.Assoc(o.Key(), o.Val())
	case IPersistentVector:
		if o.Count() != 2 {
			panic("Vector arg to map conj must be a pair")
		}
		return a.Assoc(o.Nth(0, nil), o.Nth(1, nil))
	}
	ret := a
	for es := RT.Seq(obj); es != nil; es = es.Next() {
		e := es.First().(IMapEntry)
		ret = ret.Assoc(e.Key(), e.Val()).(APersistentMap)
	}
	return ret
}

func APersistentMap_Equals(a APersistentMap, obj interface{}) bool {
	return MapEquals(a, obj)
}

// TODO: Make this function work with Glojure maps *and* pure Go maps.
// Right now it only works with the latter.
func MapEquals(m1 IPersistentMap, obj interface{}) bool {
	if m1 == obj {
		return true
	}
	switch m := obj.(type) {
	case map[interface{}]interface{}:
		if len(m) != m1.Count() {
			return false
		}
		return false
	case IPersistentMap:
		for s := m1.Seq(); s != nil; s = s.Next() {
			e := s.First().(IMapEntry)

			found := m.ContainsKey(e.Key())

			if !found || !Util.Equals(e.Val(), m.ValAt(e.Key(), nil)) {
				return false
			}
		}
	// TODO: have a map proxy for Map types in Glojure
	}

	return true
}

// TODO
func APersistentMap_Equiv(a APersistentMap, o interface{}) bool {
	return true
}

// TODO: Not sure how to figure this out. Maybe cast to an abstract hashable struct?
func APersistentMap_HashCode(a APersistentMap) int {
	/*
	if a._hash == -1 {
		a._hash = MapHash(a)
	}
	return a._hash
	*/
	return 0
}

// TODO
func MapHash(m IPersistentMap) int {
	return 0
}

// TODO: See note on APersistentMap_HashCode
func APersistentMap_HashEq(a APersistentMap) int {
	/*
	if a._hasheq == -1 {
		a._hasheq = HashUnordered(a)
	}
	return a._hasheq
	*/
	return 0
}

func MapHashEq(m IPersistentMap) int {
	return HashUnordered(m)
}

func APersistentMap_Invoke(a APersistentMap, key interface{}, notFound interface{}) interface{} {
	return a.ValAt(key, notFound)
}

/*
	Expected interface of java.util.Map.

	We can't do something similar in Go, so I'm still figuring this out.

	This is a TODO.
*/

func APersistentMap_Clear(a APersistentMap) {
	panic(UnsupportedOperationException)
}

// TODO
func APersistentMap_ContainsValue(a APersistentMap, val interface{}) bool {
	return APersistentMap_Values(a).ContainsKey(val)
}

func APersistentMap_Get(a APersistentMap, key interface{}) interface{} {
	return a.ValAt(key, nil)
}

func APersistentMap_IsEmpty(a APersistentMap) bool {
	return a.Count() == 0
}

// NOTE: In Java, this returns a set primitive. Go doesn't have these, so
// we return an IPersistentSet.
func APersistentMap_EntrySet(a APersistentMap) IPersistentSet {
	return nil
}

// NOTE: In Java, this returns a set primitive. Go doesn't have these, so
// we return an IPersistentSet.
func APersistentMap_KeySet(a APersistentMap) IPersistentSet {
	return nil
}

// Assoc in a new value.
func APersistentMap_Put(a APersistentMap, k interface{}, v interface{}) interface{} {
	panic(UnsupportedOperationException)
}

// Take another map and merge it in.
func APersistentMap_PutAll(a APersistentMap, m interface{}) {
	panic(UnsupportedOperationException)
}

func APersistentMap_Remove(a APersistentMap, key interface{}) {
	panic(UnsupportedOperationException)
}

func APersistentMap_Size(a APersistentMap) int {
	return a.Count()
}

// TODO
func APersistentMap_Values(a APersistentMap) IPersistentVector {
	return nil
}

/*
	APersistentMap required interface methods
*/

func APersistentMap_Assoc(a APersistentMap, k interface{}, v interface{}) Associative {
	panic(AbstractClassMethodException)
}

func APersistentMap_AssocEx(a APersistentMap, k interface{}, v interface{}) IPersistentMap {
	panic(AbstractClassMethodException)
}

func APersistentMap_ContainsKey(a APersistentMap, k interface{}) bool {
	panic(AbstractClassMethodException)
}

func APersistentMap_Count(a APersistentMap) int {
	panic(AbstractClassMethodException)
}

func APersistentMap_Empty(a APersistentMap) IPersistentCollection {
	panic(AbstractClassMethodException)
}

func APersistentMap_EntryAt(a APersistentMap, i interface{}) IMapEntry {
	panic(AbstractClassMethodException)
}

func APersistentMap_Seq(a APersistentMap) ISeq {
	panic(AbstractClassMethodException)
}

func APersistentMap_ValAt(a APersistentMap, key interface{}, notFound interface{}) interface{} {
	panic(AbstractClassMethodException)
}

func APersistentMap_Without(a APersistentMap, key interface{}) IPersistentMap {
	panic(AbstractClassMethodException)
}

/*
	KeySeq

	Implements abstract class ASeq
*/

type KeySeq struct {
	_meta    IPersistentMap
	seq      ISeq
	iterable Iterable // TODO: fuck
}

func CreateKeySeq(seq ISeq) *KeySeq {
	if seq == nil {
		return nil
	}
	return &KeySeq{
		seq:      seq,
		iterable: nil,
	}
}

func CreateKeySeqFromMap(m IPersistentMap) *KeySeq {
	if m == nil {
		return nil
	}
	seq := m.Seq()
	if seq == nil {
		return nil
	}
	return &KeySeq{
		seq:      seq,
		iterable: m,
	}
}

func (k *KeySeq) First() interface{} {
	e := k.seq.First().(IMapEntry)
	return e.Key()
}

func (k *KeySeq) Next() ISeq {
	return CreateKeySeq(k.seq.Next())
}

func (k *KeySeq) WithMeta(meta IPersistentMap) *KeySeq {
	return &KeySeq{
		_meta:    meta,
		seq:      k.seq,
		iterable: k.iterable,
	}
}

// TODO: No idea how to deal with this.
func (k *KeySeq) Iterator() *Iterator {
	return nil
}

/*
	Abstract methods (KeySeq)
 */

func (k *KeySeq) Cons(i interface{}) IPersistentCollection {
	return ASeq_Cons(k, i)
}

func (k *KeySeq) Count() int {
	return ASeq_Count(k)
}

func (k *KeySeq) Empty() IPersistentCollection {
	return ASeq_Empty(k)
}

func (k *KeySeq) Equals(o interface{}) bool {
	return ASeq_Equals(k, o)
}

func (k *KeySeq) Equiv(o interface{}) bool {
	return ASeq_Equiv(k, o)
}

func (k *KeySeq) HashCode() int {
	return ASeq_HashCode(k)
}

func (k *KeySeq) HashEq() int {
	return ASeq_HashEq(k)
}

func (k *KeySeq) More() ISeq {
	return ASeq_More(k)
}

func (k *KeySeq) Seq() ISeq {
	return ASeq_Seq(k)
}

func (k *KeySeq) String() string {
	return ASeq_String(k)
}


/*
	ValSeq
*/

type ValSeq struct {
	ASeq

	_meta    IPersistentMap
	seq      ISeq
	iterable Iterable
}

func CreateValSeq(seq ISeq) *ValSeq {
	if seq == nil {
		return nil
	}
	return &ValSeq{
		seq:      seq,
		iterable: nil,
	}
}

func CreateValSeqFromMap(m IPersistentMap) *ValSeq {
	if m == nil {
		return nil
	}
	seq := m.Seq()
	if seq == nil {
		return nil
	}
	return &ValSeq{
		seq:      seq,
		iterable: m,
	}
}

func (v *ValSeq) First() interface{} {
	e := v.seq.First().(IMapEntry)
	return e.Key()
}

func (v *ValSeq) Next() ISeq {
	return CreateValSeq(v.seq.Next())
}

func (v *ValSeq) WithMeta(meta IPersistentMap) *ValSeq {
	return &ValSeq{
		_meta:    meta,
		seq:      v.seq,
		iterable: v.iterable,
	}
}

// TODO: No idea how to deal with this.
func (v *ValSeq) Iterator() *Iterator {
	return nil
}

/*
	Small private anonymous classes
*/

type make_entry struct {
	AFn
}

func (m *make_entry) Invoke(key interface{}, val interface{}) interface{} {
	return CreateMapEntry(key, val)
}

type make_key struct {
	AFn
}

func (m *make_key) Invoke(key interface{}, val interface{}) interface{} {
	return key
}

type make_val struct {
	AFn
}

func (m *make_val) Invoke(key interface{}, val interface{}) interface{} {
	return val
}
