package lang
import "container/list"

/*
	PersistentHashSet

	Implements: IObj, IEditableCollection
 */

type PersistentHashSet struct {
	APersistentSet

	_meta   IPersistentMap
	_hash   int
	_hasheq int
	impl    IPersistentMap
}

var EMPTY_PERSISTENT_HASH_SET = &PersistentHashSet{
	impl: EMPTY_PERSISTENT_HASH_MAP,
}

func CreatePersistentHashSetFromInterfaceSlice(init ...interface{}) *PersistentHashSet {
	// TODO: Fix this
	/*
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	for i := 0; i < len(init); i++ {
		ret = ret.Conj(init[i]).(*TransientHashSet)
	}
	return ret.Persistent().(*PersistentHashSet)
	*/
	return nil
}

func CreatePersistentHashSetFromList(l *list.List) *PersistentHashSet {
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	for v := l.Front(); v != nil; v = v.Next() {
		ret = ret.Conj(v.Value).(*TransientHashSet)
	}
	return ret.Persistent().(*PersistentHashSet)
}

func CreatePersistentHashSetFromISeq(items ISeq) *PersistentHashSet {
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	for ; items != nil; items = items.Next() {
		ret = ret.Conj(items.First()).(*TransientHashSet)
	}
	return ret.Persistent().(*PersistentHashSet)
}

func CreatePersistentHashSetFromInterfaceSliceWithCheck(init ...interface{}) *PersistentHashSet {
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	for i := 0; i < len(init); i++ {
		ret = ret.Conj(init[i]).(*TransientHashSet)
		if ret.Count() != i + 1 {
			panic("Duplicate key: ") // + init[i]
		}
	}
	return ret.Persistent().(*PersistentHashSet)
}

func CreatePersistentHashSetFromListWithCheck(l *list.List) *PersistentHashSet {
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	i := 0
	for v := l.Front(); v != nil; v = v.Next() {
		ret = ret.Conj(v.Value).(*TransientHashSet)
		if ret.Count() != i + 1 {
			panic("Duplicate key: ") // + init[i]
		}
		i++
	}
	return ret.Persistent().(*PersistentHashSet)
}

func CreatePersistentHashSetFromISeqWithCheck(items ISeq) *PersistentHashSet {
	ret := EMPTY_PERSISTENT_HASH_SET.AsTransient()
	i := 0
	for ; items != nil; items = items.Next() {
		ret = ret.Conj(items.First()).(*TransientHashSet)
		if ret.Count() != i + 1 {
			panic("Duplicate key: ") // + init[i]
		}
		i++
	}
	return ret.Persistent().(*PersistentHashSet)
}

func (s *PersistentHashSet) Disjoin(key interface{}) IPersistentSet {
	if s.Contains(key) {
		return &PersistentHashSet{
			_meta: s.Meta(),
			impl: s.impl.Without(key),
		}
	}
	return s
}

func (s *PersistentHashSet) Cons(o interface{}) IPersistentCollection {
	if s.Contains(o) {
		return s
	}
	return &PersistentHashSet{
		_meta: s.Meta(),
		impl: s.impl.Assoc(o, o).(IPersistentMap),
	}
}

func (s *PersistentHashSet) Empty() IPersistentCollection {
	return EMPTY_PERSISTENT_HASH_SET.WithMeta(s.Meta())
}

func (s *PersistentHashSet) WithMeta(meta IPersistentMap) *PersistentHashSet {
	return &PersistentHashSet{
		_meta: meta,
		impl: s.impl,
	}
}

func (s *PersistentHashSet) AsTransient() *TransientHashSet {
	return &TransientHashSet{
		impl: s.impl.(*PersistentHashMap).AsTransient(),
	}
}

func (s *PersistentHashSet) Meta() IPersistentMap {
	return s._meta
}

type TransientHashSet struct {
	ATransientSet

	_meta IPersistentMap
	impl  ITransientMap
}

func (t *TransientHashSet) Persistent() IPersistentCollection {
	return &PersistentHashSet{
		impl: t.impl.Persistent().(IPersistentMap),
	}
}

