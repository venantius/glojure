package lang

type IChunk interface {
	Indexed

	dropFirst() IChunk
	reduce(f IFn, start interface{}) interface{}
}
