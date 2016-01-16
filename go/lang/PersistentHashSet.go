package lang

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

// TODO
func CreatePersistentHashSetFromInterfaceSlice(init ...interface{}) *PersistentHashSet {
	return nil
}

// TODO
func CreatePersistentHashSetFromList(init ...interface{}) *PersistentHashSet {
	return nil
}

// TODO
func CreatePersistentHashSetFromISeq(init ...interface{}) *PersistentHashSet {
	return nil
}

// TODO
func CreatePersistentHashSetFromInterfaceSliceWithCheck(init ...interface{}) *PersistentHashSet {
	return nil
}

// TODO
func CreatePersistentHashSetFromListWithCheck(init ...interface{}) *PersistentHashSet {
	return nil
}

// TODO
func CreatePersistentHashSetFromISeqWithCheck(init ...interface{}) *PersistentHashSet {
	return nil
}
