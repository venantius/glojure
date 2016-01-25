package lang

type ITransientAssociative interface {
	ITransientCollection
	ILookup

	Assoc(key interface{}, val interface{}) ITransientAssociative
}
