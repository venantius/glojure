package lang

type ILookupThunk interface {
	Get(target interface{}) interface{}
}
