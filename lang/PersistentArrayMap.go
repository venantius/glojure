package lang

import (
	"fmt"
)

/*
	PersistentArrayMap

	Simple implementation of persistent map on an array. Note that instances
	of this class are constant values, i.e. add/remove etc. return new
	values. Copies entire array on every change, so only appropriate for very
	small maps. Null keys and values are okay, but you won't be able to
	distinguish a null value via `ValAt` - use `Contains` or `EntryAt`

	Implements: IObj, IEditableCollection, IMapIterable, IKVReduce
*/

type PersistentArrayMap struct {
	AFn

	_meta IPersistentMap
	array []interface{}
}

const (
	HASHTABLE_THRESHOLD = 16
)

var EMPTY_PERSISTENT_ARRAY_MAP = &PersistentArrayMap{
	array: make([]interface{}, 0),
	_meta: nil,
}

func (a *PersistentArrayMap) String() string {
	return RT.PrintString(a)
}

func CreatePersistentArrayMapFromMap(other map[interface{}]interface{}) IPersistentMap {
	ret := EMPTY_PERSISTENT_ARRAY_MAP.AsTransient()
	for o := MapEntrySet(other).Seq(); o != nil; o = o.Next() {
		e := o.First().(MapEntry)
		ret = ret.Assoc(e.GetKey(), e.GetValue()).(*TransientArrayMap)
	}
	return ret.Persistent().(*PersistentArrayMap)
}

func (m *PersistentArrayMap) WithMeta(meta IPersistentMap) *PersistentArrayMap {
	return &PersistentArrayMap{
		_meta: meta,
		array: m.array,
	}
}

func (m *PersistentArrayMap) create(init []interface{}) *PersistentArrayMap {
	return &PersistentArrayMap{
		array: init,
	}
}

func (m *PersistentArrayMap) createHT(init []interface{}) IPersistentMap {
	return CreatePersistentHashMap(init).WithMeta(m.Meta())
}

func CreatePersistentArrayMapWithCheck(init []interface{}) *PersistentArrayMap {
	for i := 0; i < len(init); i += 2 {
		for j := i+2; j < len(init); j += 2 {
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

func (m *PersistentArrayMap) Count() int {
	return len(m.array) / 2
}

func (m *PersistentArrayMap) ContainsKey(key interface{}) bool {
	return m.indexOf(key) >= 0
}

func (m *PersistentArrayMap) EntryAt(key interface{}) IMapEntry {
	i := m.indexOf(key)
	if i >= 0 {
		return CreateMapEntry(m.array[i], m.array[i+1])
	}
	return nil
}

func (m *PersistentArrayMap) AssocEx(key interface{}, val interface{}) IPersistentMap {
	i := m.indexOf(key)
	var newArray []interface{}
	if i >= 0 {
		panic("Key already present.")
	} else { // didn't have key, grow
		if len(m.array) > HASHTABLE_THRESHOLD {
			return m.createHT(m.array).AssocEx(key, val)
		}
		newArray = make([]interface{}, len(m.array)+2)
		if len(m.array) > 0 {
			copy(newArray[:len(m.array)], m.array[2:len(m.array)])
		}
		newArray[0] = key
		newArray[1] = val
	}
	return m.create(newArray)
}

func (m *PersistentArrayMap) Assoc(key interface{}, val interface{}) Associative {
	i := m.indexOf(key)
	var newArray []interface{}
	if i >= 0 { // already have key, same-sized replacement
		if m.array[i+1] == val { // no change, no op
			return m
		}
		copy(newArray, m.array)
		newArray[i+1] = val
	} else { // didn't have key, grow
		if len(m.array) > HASHTABLE_THRESHOLD {
			return m.createHT(m.array).Assoc(key, val)
		}
		newArray = make([]interface{}, len(m.array)+2)
		if len(m.array) > 0 {
			copy(newArray, m.array)
		}
		newArray[len(newArray)-2] = key
		newArray[len(newArray)-1] = val
	}
	return m.create(newArray)
}

func (m *PersistentArrayMap) Without(key interface{}) IPersistentMap {
	i := m.indexOf(key)
	if i >= 0 { // have key, will remove
		newlen := len(m.array) - 2
		if newlen == 0 {
			return m.Empty().(IPersistentMap)
		}
		newArray := make([]interface{}, newlen)
		copy(newArray[0:i], m.array[0:i])
		copy(newArray[i+2:newlen-i], m.array[i+2:newlen-i])
		return m.create(newArray)
	}
	// don't have key, no operation
	return m
}

func (m *PersistentArrayMap) Empty() IPersistentCollection {
	return EMPTY_PERSISTENT_ARRAY_MAP.WithMeta(m.Meta())
}

func (m *PersistentArrayMap) ValAt(key interface{}, notFound interface{}) interface{} {
	i := m.indexOf(key)
	if i >= 0 {
		return m.array[i+1]
	}
	return notFound
}

// TODO: Why is this necessary?
func (m *PersistentArrayMap) Capacity() int {
	return m.Count()
}

func (m *PersistentArrayMap) indexOfObject(key interface{}) int {
	ep := Util.EquivPred(key)
	for i := 0; i < len(m.array); i += 2 {
		if ep.Equiv(key, m.array[i]) {
			return i
		}
	}
	return -1
}

func (m *PersistentArrayMap) indexOf(key interface{}) int {
	switch k := key.(type) {
	case Keyword:
		for i := 0; i < len(m.array); i += 2 {
			if k == m.array[i] {
				return i
			}
		}
		return -1
	}
	return m.indexOfObject(key)
}

func equalKey(k1 interface{}, k2 interface{}) bool {
	switch k1.(type) {
	case Keyword:
		return k1 == k2
	}
	return Util.Equiv(k1, k2)
}

// TODO: As always, not sure about this
func (m *PersistentArrayMap) Iterator() *Iterator {
	return nil
}

// TODO
func (m *PersistentArrayMap) KeyIterator() *Iterator {
	return nil
}

// TODO
func (m *PersistentArrayMap) ValIterator() *Iterator {
	return nil
}

func (m *PersistentArrayMap) Seq() ISeq {
	if len(m.array) > 0 {
		return &PersistentArrayMapSeq{
			array: m.array,
			i:     0,
		}
	}
	return nil
}

func (m *PersistentArrayMap) Meta() IPersistentMap {
	return m._meta
}

func (m *PersistentArrayMap) KVReduce(f IFn, init interface{}) interface{} {
	for i := 0; i < len(m.array); i += 2 {
		init = f.Invoke(init, m.array[i], m.array[i+1])
		if RT.IsReduced(init) {
			return init.(IDeref).Deref()
		}
	}
	return init
}

func (m *PersistentArrayMap) AsTransient() *TransientArrayMap {
	newArr := make([]interface{}, len(m.array))
	copy(newArr, m.array)
	return &TransientArrayMap{
		edit:  true,
		array: newArr,
		len:   len(m.array),
	}
}

/*
	Abstract methods (PersistentArrayMap)
 */

func (m *PersistentArrayMap) Cons(o interface{}) IPersistentCollection{
	return APersistentMap_Cons(m, o)
}

func (m *PersistentArrayMap) Equals(o interface{}) bool {
	return APersistentMap_Equals(m, o)
}

func (m *PersistentArrayMap) Equiv(o interface{}) bool{
	return APersistentMap_Equiv(m, o)
}

func (m *PersistentArrayMap) HashEq() int{
	return APersistentMap_HashEq(m)
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

func (ms *PersistentArrayMapSeq) First() interface{} {
	return CreateMapEntry(ms.array[ms.i], ms.array[ms.i+1])
}

func (ms *PersistentArrayMapSeq) Next() ISeq {
	if (ms.i + 2) < len(ms.array) {
		return &PersistentArrayMapSeq{
			array: ms.array,
			i:     ms.i + 2,
		}
	}
	return nil
}

func (ms *PersistentArrayMapSeq) Count() int {
	return (len(ms.array) - ms.i) / 2
}

func (ms *PersistentArrayMapSeq) WithMeta(meta IPersistentMap) interface{} {
	return &PersistentArrayMapSeq{
		_meta: meta,
		array: ms.array,
		i:     ms.i,
	}
}

/*
	TransientArrayMap

	Implements abstract class ATransientArrayMap
*/

type TransientArrayMap struct {
	AFn

	// NOTE: in JVM Clojure, we also have `volatile Thread owner`
	// We've changed that to `edit` here, as with TransientVectors
	edit  bool
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

func (t *TransientArrayMap) doAssoc(key interface{}, val interface{}) ITransientMap {
	i := t.indexOf(key)
	if i >= 0 { // already have key
		if t.array[i+1] != val { // no change, no op
			t.array[i+1] = val
		}
	} else {
		if t.len >= len(t.array) {
			return CreatePersistentHashMap(t.array).AsTransient().Assoc(key, val)
		}

		t.len++
		t.array[t.len] = key

		t.len++
		t.array[t.len] = val

	}
	return t
}

func (t *TransientArrayMap) doWithout(key interface{}) ITransientMap {
	i := t.indexOf(key)
	if i >= 0 { // have key, will remove
		if t.len >= 2 {
			t.array[i] = t.array[t.len-2]
			t.array[i+1] = t.array[t.len-1]
		}
		t.len -= 2
	}
	return t
}

func (t *TransientArrayMap) doValAt(key interface{}, notFound interface{}) interface{} {
	i := t.indexOf(key)
	if i >= 0 {
		return t.array[i+1]
	}
	return notFound
}

func (t *TransientArrayMap) doCount() int {
	return t.len / 2
}

func (t *TransientArrayMap) doPersistent() IPersistentMap {
	t.ensureEditable()
	t.edit = false
	a := make([]interface{}, t.len)
	copy(a[:t.len], t.array[:t.len])
	return &PersistentArrayMap{
		array: a,
	}
}

func (t *TransientArrayMap) ensureEditable() {
	if t.edit == false {
		panic(TransientUsedAfterPersistentCallError)
	}
}

/*
	Abstract methods (TransientArrayMap)
 */

func (t *TransientArrayMap) Assoc(k interface{}, v interface{}) ITransientMap {
	return ATransientMap_Assoc(t, k, v)
}

func (t *TransientArrayMap) Conj(o interface{}) ITransientCollection {
	return ATransientMap_Conj(t, o)
}

func (t *TransientArrayMap) Count() int {
	return ATransientMap_Count(t)
}

func (t *TransientArrayMap) Persistent() IPersistentCollection {
	return ATransientMap_Persistent(t)
}

func (t *TransientArrayMap) ValAt(key interface{}, notFound interface{}) interface{} {
	return ATransientMap_ValAt(t, key, notFound)
}

func (t *TransientArrayMap) Without(key interface{}) ITransientMap {
	return ATransientMap_Without(t, key)
}
