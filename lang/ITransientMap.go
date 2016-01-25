package lang

type ITransientMap interface {
	ITransientAssociative
	Counted

	// Assoc(key interface{}, val interface{}) ITransientMap
	Without(key interface{}) ITransientMap
	// Persistent() IPersistentMap
}
