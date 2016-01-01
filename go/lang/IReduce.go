package lang

type IReduce interface {
	IReduceInit

	Reduce(f IFn) interface{}
}
