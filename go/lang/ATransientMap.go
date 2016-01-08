package lang

// NOTE: Implements ITransientMap
type ATransientMap struct {
	AFn
}

// TODO
func (t *ATransientMap) Conj(o interface{}) ITransientCollection {
	panic(NotYetImplementedException)
}

func (t *ATransientMap) Invoke(arg1 interface{}, notFound interface{}) interface{} {
	return t.ValAt(arg1, notFound)
}

func (t *ATransientMap) Assoc(key interface{}, val interface{}) ITransientAssociative {
	t.ensureEditable()
	return t.doAssoc(key, val)
}

func (t *ATransientMap) Without(key interface{}) ITransientMap {
	t.ensureEditable()
	return t.doWithout(key)
}

func (t *ATransientMap) Persistent() IPersistentCollection {
	t.ensureEditable()
	return t.doPersistent()
}

func (t *ATransientMap) ValAt(key interface{}, notFound interface{}) interface{} {
	t.ensureEditable()
	return t.doValAt(key, notFound)
}

func (t *ATransientMap) Count() int {
	t.ensureEditable()
	return t.doCount()
}

/*
	Abstract methods
*/

func (t *ATransientMap) doAssoc(key interface{}, val interface{}) ITransientMap {
	panic(AbstractClassMethodException)
}

func (t *ATransientMap) doCount() int {
	panic(AbstractClassMethodException)
}

func (t *ATransientMap) doPersistent() IPersistentCollection {
	panic(AbstractClassMethodException)
}
func (t *ATransientMap) doValAt(k interface{}, notFound interface{}) interface{} {
	panic(AbstractClassMethodException)
}

func (t *ATransientMap) doWithout(key interface{}) ITransientMap {
	panic(AbstractClassMethodException)
}

func (t *ATransientMap) ensureEditable() {
	panic(AbstractClassMethodException)
}
