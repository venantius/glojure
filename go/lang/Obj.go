package lang

// NOTE: Abstract
// NOTE: Implements IObj, Serializable
type Obj struct {
	_meta IPersistentMap
}

func (o *Obj) Meta() IPersistentMap {
	return o._meta
}

func (o *Obj) WithMeta(meta IPersistentMap) interface{} {
	// TODO: use central exception
	panic("Not implemented!")
}
