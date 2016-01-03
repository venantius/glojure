package lang

type IPersistentVector interface {
	Associative
	Sequential
	IPersistentStack
	Reversible
	Indexed

	Length() int
	AssocN(i int, val interface{}) IPersistentVector
	Cons(i interface{}) IPersistentVector
}
