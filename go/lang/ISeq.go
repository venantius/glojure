package lang

type ISeq interface {
	first() interface{}
	next() ISeq
	more() ISeq
	cons(interface{}) ISeq
}
