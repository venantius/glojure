package lang

type ILookup interface {
	valAt(key interface{}, notFound interface{}) interface{}
}
