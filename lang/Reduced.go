package lang

// NOTE: Implements IDeref
type Reduced struct {
	val interface{}
}

func (r *Reduced) Deref() interface{} {
	return r.val
}
