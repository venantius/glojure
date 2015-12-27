package lang

type ISeq interface {
	First() interface{}
	Next() ISeq
	More() ISeq
	Cons(interface{}) ISeq
}
