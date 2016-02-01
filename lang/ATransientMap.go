package lang

/*
	ATransientMap

	Implements: ITransientMap
*/

type ATransientMap interface {
	ITransientMap
	IFn

	ensureEditable()
	doAssoc(key interface{}, val interface{}) ITransientMap
	doWithout(key interface{}) ITransientMap
	doValAt(key interface{}, notFound interface{}) interface{}
	doCount() int
	doPersistent() IPersistentMap
}

func ATransientMap_Conj(t ATransientMap, o interface{}) ITransientCollection {
	t.ensureEditable()
	switch obj := o.(type) {
	case MapEntry:
		return t.Assoc(obj.GetKey(), obj.GetValue())
	case IPersistentVector:
		if obj.Count() != 2 {
			panic("Vector arg to map conj must be a pair")
		}
		return t.Assoc(obj.Nth(0, nil), obj.Nth(1, nil))
	}
	ret := t
	for es := RT.Seq(o); es != nil; es = es.Next() {
		var e MapEntry = es.First().(MapEntry)
		ret = ret.Assoc(e.GetKey(), e.GetValue()).(ATransientMap)
	}
	return ret
}

func ATransientMap_Invoke(t ATransientMap, arg1 interface{}, notFound interface{}) interface{} {
	return t.ValAt(arg1, notFound)
}

func ATransientMap_Assoc(t ATransientMap, key interface{}, val interface{}) ITransientMap {
	t.ensureEditable()
	return t.doAssoc(key, val)
}

func ATransientMap_Without(t ATransientMap, key interface{}) ITransientMap {
	t.ensureEditable()
	return t.doWithout(key)
}

func ATransientMap_Persistent(t ATransientMap) IPersistentCollection {
	t.ensureEditable()
	return t.doPersistent()
}

func ATransientMap_ValAt(t ATransientMap, key interface{}, notFound interface{}) interface{} {
	t.ensureEditable()
	return t.doValAt(key, notFound)
}

func ATransientMap_Count(t ATransientMap) int {
	t.ensureEditable()
	return t.doCount()
}

/*
	Abstract methods
*/

/*
func (t *ATransientMap) ATransientMap_doAssoc(key interface{}, val interface{}) ITransientMap {
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
*/