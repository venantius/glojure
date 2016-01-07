package lang

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

	array []interface{}
}
