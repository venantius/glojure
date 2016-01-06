package lang

type IObj interface {
	IMeta

	WithMeta(meta IPersistentMap) interface{}
	Equals(other interface{}) bool
}
