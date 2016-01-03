package lang

type ILookup interface {
	ValAt(key interface{}, notFound interface{}) interface{}
}
