package lang

type IChunk interface {
	Indexed

	DropFirst() IChunk
	Reduce(f IFn, start interface{}) interface{}
}
