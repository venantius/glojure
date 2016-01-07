package lang

type Settable interface {
	DoSet(val interface{}) interface{}
	DoReset(val interface{}) interface{}
}
