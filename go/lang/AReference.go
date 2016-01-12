package lang

/*
	AReference

	Implements: IReference
*/

type AReference struct {
	_meta IPersistentMap
}

func (a *AReference) AlterMeta(alter IFn, args ISeq) IPersistentMap {
	c := &Cons{
		_meta: a._meta,
		_more: args,
	}
	a._meta = alter.ApplyTo(c).(IPersistentMap)
	return a._meta
}

func (a *AReference) ResetMeta(m IPersistentMap) IPersistentMap {
	a._meta = m
	return m
}
