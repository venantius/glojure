package lang

// NOTE: Implements IMapEntry
type AMapEntry struct {
	APersistentVector
}

func (a *AMapEntry) Nth(i int) interface{} {
	if i == 0 {
		return a.Key()
	} else if i == 1 {
		return a.Val()
	} else {
		panic(IndexOutOfBoundsException)
	}
}

// TODO: Haven't gotten around to this at all
func (a *AMapEntry) asVector() IPersistentVector {
	return CreateOwningLazilyPersistentVector(a.Key(), a.Val())
}

func (a *AMapEntry) AssocN(i int, val interface{}) IPersistentVector {
	return a.asVector().AssocN(i, val)
}

func (a *AMapEntry) Count() int {
	return 2
}

func (a *AMapEntry) Seq() ISeq {
	return a.asVector().Seq()
}

func (a *AMapEntry) Cons(o interface{}) IPersistentCollection {
	return a.asVector().Cons(o)
}

func (a *AMapEntry) Empty() IPersistentCollection {
	return nil
}

// TODO: again with the lazily persistent vectors
func (a *AMapEntry) Pop() IPersistentStack {
	return CreateOwningLazilyPersistentVector(a.Key())
}

func (a *AMapEntry) SetValue(value interface{}) interface{} {
	panic(UnsupportedOperationException)
}

/*
	Abstract class methods
*/

func (a *AMapEntry) Key() interface{} {
	panic(AbstractClassMethodException)
}

func (a *AMapEntry) Val() interface{} {
	panic(AbstractClassMethodException)
}
