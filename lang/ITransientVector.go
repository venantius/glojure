package lang

type ITransientVector interface {
	ITransientAssociative
	Indexed

	Assoc(key interface{}, val interface{}) ITransientVector
	AssocN(i int, val interface{}) ITransientVector
	Pop() ITransientVector
}
