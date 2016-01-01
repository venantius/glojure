package lang

// TODO: Is this the right way to do this?
type IReduceInit interface {
	ReduceWithInit(f IFn, start interface{}) interface{}
}
