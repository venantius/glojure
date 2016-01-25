package lang

type IKVReduce interface {
	KVReduce(f IFn, init interface{}) interface{}
}
