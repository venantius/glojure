package lang

// NOTE: Implements IRef
type ARef struct {
	// TODO: Inherit AReference fields
	*AReference

	_meta     IPersistentMap
	validator IFn
	watches   IPersistentMap
}

// TODO
func (a *ARef) validate(vf IFn, val interface{}) {
	// try catch
}

func (a *ARef) SetValidator(vf IFn) {
	a.validate(vf, a.Deref())
	a.validator = vf
}

func (a *ARef) GetValidator() IFn {
	return a.validator
}

func (a *ARef) GetWatches() IPersistentMap {
	return a.watches
}

// TODO: what is keyword synchronized?
func (a *ARef) AddWatch(key interface{}, callback IFn) IRef {
	a.watches = a.watches.Assoc(key, callback).(IPersistentMap) // TODO: is this cheating?
	return a
}

func (a *ARef) RemoveWatch(key interface{}) IRef {
	a.watches = a.watches.Without(key)
	return a
}

// TODO
func (a *ARef) NotifyWatches(oldval interface{}, newval interface{}) {
	// TODO: Remove this nil check
	if a.watches != nil {
		ws := a.watches
		if ws.Count() > 0 {
			for s := ws.Seq(); s != nil; s = s.Next() {
				// NOTE this is a little different from the Java implementation
				e := s.First().(IMapEntry) // TODO: huh?
				fn := e.Val().(IFn)
				if fn != nil {
					fn.Invoke(e.Key(), a, oldval, newval)
				}
			}
		}
	}
}

/*
	The following methods must be implemented by the concrete classes
*/

func (a *ARef) Deref() interface{} {
	panic(AbstractClassMethodException)
}
