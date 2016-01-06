package lang

// May need to be refactored
type IPersistentSet interface {
	IPersistentCollection
	Counted

	Disjoin(key interface{}) IPersistentSet
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
}
