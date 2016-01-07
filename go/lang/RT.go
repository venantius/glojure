package lang

// In case it wasn't obvious (it wasn't to me), RT stands for RunTime
type rt struct{}

/*
	NOTE: I've made the design decision for now to mimic static methods as best I can,
	which in this case means creating a private class and a single public object for that class. In practice I think that RT.java is more or less a catchall for a host of static methods that could just as easily be generally pure functions.

	I'll decide whether or not I want to change this later.
*/

func (_ *rt) EMPTY_ARRAY() []interface{} {
	return make([]interface{}, 1)
}

var RT = rt{} // Mock static methods

func (_ *rt) IsReduced(r interface{}) bool {
	switch r.(type) {
	case Reduced:
		return true
	default:
		return false
	}
}

// TODO....so much

func (_ *rt) Seq(coll interface{}) ISeq {
	switch c := coll.(type) {
	case *ASeq:
		return c
	case *LazySeq:
		return c.Seq()
	}
	return RT.seqFrom(coll)
}

func (_ *rt) seqFrom(coll interface{}) ISeq {
	// TODO
	return nil
}

func (_ *rt) SubVec(v IPersistentVector, start int, end int) IPersistentVector {
	if end < start || start < 0 || end > v.Count() {
		panic(IndexOutOfBoundsException)
	}
	if start == end {
		return nil
		// return EMPTY_PERSISTENT_VECTOR TODO
	}
	return &SubVector{} // TODO
}

func (_ *rt) getFrom(coll interface{}, key interface{}, notFound interface{}) interface{} {
	if coll == nil {
		return nil
	}
	// TODO: This implementation is incomplete
	return nil
}

func (_ *rt) Get(coll interface{}, key interface{}, notFound interface{}) interface{} {
	switch coll.(type) {
	case ILookup:
		return coll.(ILookup).ValAt(key, notFound)
	}
	return RT.getFrom(coll, key, notFound)
}

// unordered

// TODO
func (_ *rt) Count(o interface{}) int {
	return 0
}

// TODO
func (_ *rt) PrintString(o interface{}) string {
	return ""
}

// TODO
func (_ *rt) ToArray(coll interface{}) []interface{} {
	return nil
}
