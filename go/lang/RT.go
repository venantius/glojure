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
	switch coll.(type) {
	case ASeq:
		return coll.(*ASeq)
	case LazySeq:
		return coll.(*LazySeq).Seq()
	}
	return RT.seqFrom(coll)
}

func (_ *rt) seqFrom(coll interface{}) ISeq {
	// TODO
	return nil
}
