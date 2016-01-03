package lang

type IAtom interface {
	Swap(args ...interface{}) interface{}
	CompareAndSet(oldv interface{}, newv interface{}) bool
	Reset(newval interface{}) interface{}
}
