package lang

import (
	"fmt"
)

/*
	Simple implementation of persistent map on an array. Note that instances
	of this class are constant values, i.e. add/remove etc. return new
	values. Copies entire array on every change, so only appropriate for very
	small maps. Null keys and values are okay, but you won't be able to
	distinguish a null value via `ValAt` - use `Contains` or `EntryAt`
*/

// NOTE: Implements IObj, IEditableCollection, IMapIterable, IKVReduce
type PersistentArrayMap struct {
	APersistentMap

	_meta IPersistentMap
	array []interface{}
}

const (
	HASHABLE_THRESHOLD = 16
)

var EMPTY_PERSISTENT_ARRAY_MAP = PersistentArrayMap{
	array: make([]interface{}, HASHABLE_THRESHOLD), // NOTE: the length might be wrong.
	_meta: nil,
}

// TODO: Rewrite me!
// This probably needs a utility function that returns an iterator.
func CreatePersistentArrayMapFromMap(other map[interface{}]interface{}) IPersistentMap {
	ret := EMPTY_PERSISTENT_ARRAY_MAP.AsTransient()
	for o := other.EntrySet().Seq(); o != nil; o = o.Next() {
		e := o.(IMapEntry)
		ret = ret.Assoc(o.GetKey(), o.GetValue())
	}
	return ret.Persistent()
}

func (m *PersistentArrayMap) WithMeta(meta IPersistentMap) *PersistentArrayMap {
	return &PersistentArrayMap{
		_meta: meta,
		array: m.array,
	}
}

func (m *PersistentArrayMap) create(init ...interface{}) *PersistentArrayMap {
	return &PersistentArrayMap{
		array: init,
	}
}

func (m *PersistentArrayMap) createHT(init []interface{}) IPersistentMap {
	return PersistentHashMap.Create(m.Meta(), init)
}

func CreatePersistentArrayMapWithCheck(init []interface{}) *PersistentArrayMap {
	for i := 0; i < len(init); i += 2 {
		for j := 0; j < len(init); j += 2 {
			if equalKey(init[i], init[j]) {
				panic(DuplicateKeyException(init[i].(fmt.Stringer)))
			}
		}
	}
	return &PersistentArrayMap{
		array: init,
	}
}

func CreateAsIfByAssoc(init []interface{}) *PersistentArrayMap {
	if len(init)&1 == 1 {
		panic(NoValueSuppliedForKeyException(init[len(init)-1].(fmt.Stringer)))
	}

	// NOTE: If this looks like it is doing busy-work, it is because it
	// is achieving these goals: O(n^2) runtime like CreateWithCheck(),
	// never modify init arg, and only allocate memory if there are
	// duplicate keys.
	n := 0
	for i := 0; i < len(init); i += 2 {
		duplicateKey := false
		for j := 0; j < i; j += 2 {
			if equalKey(init[i], init[j]) {
				duplicateKey = true
				break
			}
		}
		if !duplicateKey {
			n += 2
		}
	}
	if n < len(init) {
		nodups := make([]interface{}, n)
		m := 0
		for i := 0; i < len(init); i += 2 {
			duplicateKey := false
			for j := 0; j < m; j += 2 {
				if equalKey(init[i], nodups[j]) {
					duplicateKey = true
					break
				}
			}
			if !duplicateKey {
				var j int
				for j = len(init) - 2; j >= i; j -= 2 {
					if equalKey(init[i], init[j]) {
						break
					}
				}
				nodups[m] = init[i]
				nodups[m+1] = init[j+1]
				m += 2
			}
		}
		if m != n {
			panic(IllegalArgumentException)
		}
		init = nodups
	}
	return &PersistentArrayMap{
		array: init,
	}
}

func (p *PersistentArrayMap) Count() int {
	return len(p.array) / 2
}

func (p *PersistentArrayMap) ContainsKey(key interface{}) bool {
	return p.indexOf(key) >= 0
}

func (p *PersistentArrayMap) EntryAt(key interface{}) IMapEntry {
	i := p.indexOf(key)
	if i >= 0 {
		return CreateMapEntry(p.array[i], p.array[i+1])
	}
	return nil
}

func (p *PersistentArrayMap) AssocEx(key interface{}, val interface{}) IPersistentMap {
	i := p.indexOf(key)
	var newArray []interface{}
	if i >= 0 {
		panic("Key already present.")
	} else { // didn't have key, grow
		if len(p.array) > HASHABLE_THRESHOLD {
			return p.createHT(p.array).AssocEx(key, val)
		}
		newArray := make([]interface{}, len(p.array)+2)
		if len(p.array) > 0 {
			copy(newArray, p.array)
		}
		newArray[0] = key
		newArray[1] = val
	}
	return create(newArray)
}

// TODO
func (p *PersistentArrayMap) Assoc(key interface{}, val interface{}) IPersistentMap {
	return nil
}

// TODO
func (p *PersistentArrayMap) Without(key interface{}) IPersistentMap {
	return nil
}

// TODO
func (p *PersistentArrayMap) Empty() IPersistentMap {
	return EMPTY_PERSISTENT_MAP.WithMeta(p.Meta())
}

// TODO
func (p *PersistentArrayMap) ValAt(key interface{}, notFound interface{}) interface{} {
	return nil
}

// TODO: Why is this necessary?
func (p *PersistentArrayMap) Capacity() int {
	return p.Count()
}

// TODO
func (p *PersistentArrayMap) indexOfObject(key interface{}) int {
	return 0
}

// TODO
func (p *PersistentArrayMap) indexOf(key interface{}) int {
	return 1
}

func equalKey(k1 interface{}, k2 interface{}) bool {
	switch k1.(type) {
	case Keyword:
		return k1 == k2
	}
	return Util.Equiv(k1, k2)
}

// TODO: As always, not sure about this
func (p *PersistentArrayMap) Iterator() *Iterator {
	return nil
}

// TODO
func (p *PersistentArrayMap) KeyIterator() *Iterator {
	return nil
}

// TODO
func (p *PersistentArrayMap) ValIterator() *Iterator {
	return nil
}

// TODO
func (p *PersistentArrayMap) Seq() ISeq {
	return nil
}

func (p *PersistentArrayMap) Meta() IPersistentMap {
	return p._meta
}

func (p *PersistentArrayMap) KVReduce(f IFn, init interface{}) interface{} {
	for i := 0; i < len(p.array); i += 2 {
		init = f.Invoke(init, p.array[i], p.array[i+1])
		if RT.IsReduced(init) {
			return init.(IDeref).Deref()
		}
	}
	return init
}

func (p *PersistentArrayMap) AsTransient() ITransientMap {
	newArr := make([]interface{}, len(p.array))
	copy(newArr, p.array)
	return &TransientArrayMap{
		array: newArr,
		len:   len(p.array),
	}
}

/*
	PersistentArrayMap.Seq
*/

// NOTE: In JVM Clojure, this is a nested class: PersistentArrayMap.Seq
type PersistentArrayMapSeq struct {
	ASeq

	_meta IPersistentMap
	array []interface{}
	i     int
}

func (ps *PersistentArrayMapSeq) First() interface{} {
	return CreateMapEntry(ps.array[ps.i], ps.array[ps.i+1])
}

func (ps *PersistentArrayMapSeq) Next() ISeq {
	if (ps.i + 2) < len(ps.array) {
		return &PersistentArrayMapSeq{
			array: ps.array,
			i:     ps.i + 2,
		}
	}
	return nil
}

func (ps *PersistentArrayMapSeq) Count() int {
	return (len(ps.array) - ps.i) / 2
}

func (ps *PersistentArrayMapSeq) WithMeta(meta IPersistentMap) interface{} {
	return &PersistentArrayMapSeq{
		_meta: meta,
		array: ps.array,
		i:     ps.i,
	}
}

/*
	TransientArrayMap
*/

type TransientArrayMap struct {
	ATransientMap

	len   int
	array []interface{}
}

func (t *TransientArrayMap) indexOf(key interface{}) int {
	for i := 0; i < t.len; i += 2 {
		if equalKey(t.array[i], key) {
			return i
		}
	}
	return -1
}

func (t *TransientArrayMap) doAssoc(key interface{}, val interface{}) ITransientAssociative {
	i := t.indexOf(key)
	if i >= 0 { // already have key
		if t.array[i+1] != val { // no change, no op
			t.array[i+1] = val
		}
	} else {
		if t.len >= len(t.array) {
			// TODO
		}

		t.len++
		t.array[t.len] = key

		t.len++
		t.array[t.len] = val

	}
	return *t
}

// TODO
func (t *TransientArrayMap) doWithout(key interface{}) ITransientMap {
	return nil
}

// TODO
func (t *TransientArrayMap) doValAt(key interface{}, notFound interface{}) interface{} {
	return nil
}

func (t *TransientArrayMap) doCount() int {
	return t.len / 2
}

// TODO
func (t *TransientArrayMap) doPersistent() IPersistentMap {
	return nil
}

// TODO
func (t *TransientArrayMap) ensureEditable() {
}
