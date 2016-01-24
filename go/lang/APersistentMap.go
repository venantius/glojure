package lang

// NOTE: Implements IPersistentMap, Map, Iterable, Serializable, MapEquivalance, IHashEq
type APersistentMap struct {
	AFn
	_hash   int
	_hasheq int
}

func (a *APersistentMap) Cons(obj interface{}) IPersistentCollection {
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
		ret = ret.Assoc(e.Key(), e.Val()).(*APersistentMap)
	}
	return ret
}

func (a *APersistentMap) Equals(obj interface{}) bool {
	return MapEquals(a, obj)
}

// TODO: Make this function work with Glojure maps *and* pure Go maps.
// Right now it only works with the latter.
func MapEquals(m1 IPersistentMap, obj interface{}) bool {
	if m1 == obj {
		return true
	}
	switch obj.(type) {
	case map[interface{}]interface{}:
	// TODO: have a map proxy for Map types in Glojure
	default:
		return false
	}
	m := obj.(map[interface{}]interface{})
	if len(m) != m1.Count() {
		return false
	}

	for s := m1.Seq(); s != nil; s = s.Next() {
		e := s.First().(IMapEntry)
		_, found := m[e.Key()]

		if !found || Util.Equals(e.Val(), m[e.Key()]) {
			return false
		}
	}
	return true
}

// TODO
func (a *APersistentMap) Equiv(o interface{}) bool {
	return true
}

func (a *APersistentMap) HashCode() int {
	if a._hash == -1 {
		a._hash = MapHash(a)
	}
	return a._hash
}

// TODO
func MapHash(m IPersistentMap) int {
	return 0
}

func (a *APersistentMap) HashEq() int {
	if a._hasheq == -1 {
		a._hasheq = HashUnordered(a)
	}
	return a._hasheq
}

func MapHashEq(m IPersistentMap) int {
	return HashUnordered(m)
}

func (a *APersistentMap) Invoke(key interface{}, notFound interface{}) interface{} {
	return a.ValAt(key, notFound)
}

/*
	Expected interface of java.util.Map.

	We can't do something similar in Go, so I'm still figuring this out.

	This is a TODO.
*/

func (a *APersistentMap) Clear() {
	panic(UnsupportedOperationException)
}

// TODO
func (a *APersistentMap) ContainsValue(val interface{}) bool {
	return a.Values().ContainsKey(val)
}

func (a *APersistentMap) Get(key interface{}) interface{} {
	return a.ValAt(key, nil)
}

func (a *APersistentMap) IsEmpty() bool {
	return a.Count() == 0
}

// NOTE: In Java, this returns a set primitive. Go doesn't have these, so
// we return an IPersistentSet.
func (a *APersistentMap) EntrySet() IPersistentSet {
	return nil
}

// NOTE: In Java, this returns a set primitive. Go doesn't have these, so
// we return an IPersistentSet.
func (a *APersistentMap) KeySet() IPersistentSet {
	return nil
}

// Assoc in a new value.
func (a *APersistentMap) Put(k interface{}, v interface{}) interface{} {
	panic(UnsupportedOperationException)
}

// Take another map and merge it in.
func (a *APersistentMap) PutAll(m interface{}) {
	panic(UnsupportedOperationException)
}

func (a *APersistentMap) Remove(key interface{}) {
	panic(UnsupportedOperationException)
}

func (a *APersistentMap) Size() int {
	return a.Count()
}

// TODO
func (a *APersistentMap) Values() IPersistentVector {
	return nil
}

/*
	APersistentMap required interface methods
*/

func (a *APersistentMap) Assoc(k interface{}, v interface{}) Associative {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) AssocEx(k interface{}, v interface{}) IPersistentMap {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) ContainsKey(k interface{}) bool {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) Count() int {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) Empty() IPersistentCollection {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) EntryAt(i interface{}) IMapEntry {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) Seq() ISeq {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) ValAt(key interface{}, notFound interface{}) interface{} {
	panic(AbstractClassMethodException)
}

func (a *APersistentMap) Without(key interface{}) IPersistentMap {
	panic(AbstractClassMethodException)
}

/*
	KeySeq
*/

type KeySeq struct {
	ASeq

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
