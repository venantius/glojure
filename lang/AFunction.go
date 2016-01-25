package lang

// NOTE: Implements IObj, Comparator, Fn, Serializable
type AFunction struct {
	AFn

	__methodImplCache MethodImplCache
}

func (a *AFunction) Meta() IPersistentMap {
	panic("Not implemented")
}

func (a *AFunction) WithMeta(meta IPersistentMap) IObj {
	panic("Not implemented")
	// TODO: Implement the anonymous class shenanigans here.
}

func (a *AFunction) Compare(obj1 interface{}, obj2 interface{}) int {
	panic("Not implemented")
}
