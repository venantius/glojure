package lang

import "sync"

/*
	ATransientSet

	Implements: ITransientSet
 */

type ATransientSet struct {
	AFn
	sync.Mutex

	impl ITransientMap
}

func (t *ATransientSet) Count() int {
	return t.impl.Count()
}

func (t *ATransientSet) Conj(val interface{}) ITransientCollection {
	// TODO: Fix this
	/*
	m := t.impl.Assoc(val, val).(ITransientMap)
	if m != t.impl {
		t.impl = m
	}
	return t*/
	return nil
}

func (t *ATransientSet) Contains(key interface{}) bool {
	return t != t.impl.ValAt(key, t)
}

func (t *ATransientSet) Disjoin(key interface{}) ITransientSet {
	m := t.impl.Without(key)
	if m != t.impl {
		t.impl = m
	}
	return t
}

func (t *ATransientSet) Get(key interface{}) interface{} {
	return t.impl.ValAt(key, nil)
}

func (t *ATransientSet) Invoke(args ...interface{}) interface{} {
	if len(args) == 1 {
		return t.impl.ValAt(args[0], nil)
	} else {
		return t.impl.ValAt(args[0], args[1])
	}
}

// Abstract struct methods

func (t *ATransientSet) Persistent() IPersistentCollection {
	panic(AbstractClassMethodException)
}