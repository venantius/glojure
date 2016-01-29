package lang

type IBlockingDeref interface {
	Deref(ms int64, timeoutValue interface{}) interface{}
}
