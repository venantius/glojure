package lang

type ITransientCollection interface {
	Conj(val interface{}) ITransientCollection
	Persistent() IPersistentCollection
}
