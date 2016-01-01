package lang

type ITransientVector interface {
	ITransientAssociative
	Indexed

	AssocN(i int, val interface{}) ITransientVector
	Pop() ITransientVector
}
