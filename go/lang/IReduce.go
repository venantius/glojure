package lang

type IReduce interface {
	IReduceInit

	Reduce(f IFn, init interface{}) interface{}
}
