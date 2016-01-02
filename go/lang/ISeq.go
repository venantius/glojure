package lang

type ISeq interface {
	First() interface{}
	Next() ISeq
	More() ISeq
	ConsISeq(interface{}) ISeq
}
