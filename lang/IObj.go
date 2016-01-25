package lang

type IObj interface {
	IMeta

	WithMeta(meta IPersistentMap) interface{}
}
