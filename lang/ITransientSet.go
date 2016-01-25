package lang

type ITransientSet interface {
	ITransientCollection
	Counted

	Disjoin(key interface{}) ITransientSet
	Contains(key interface{}) bool
	Get(key interface{}) interface{}
}
