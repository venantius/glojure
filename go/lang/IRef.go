package lang

type IRef interface {
	IDeref

	SetValidator(vf IFn)
	GetValidator() IFn
	GetWatches() IPersistentMap
	AddWatch(key interface{}, callback IFn) IRef
	RemoveWatch(key interface{}) IRef
}
