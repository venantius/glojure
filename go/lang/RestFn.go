package lang

// NOTE: Abstract class in Java
type RestFn struct {
	AFunction
}

func (r *RestFn) GetRequiredArity() int {
	// TODO: Factor this into a generic error we can use.
	panic("Not implemented")
}

func (r *RestFn) DoInvoke(args ...interface{}) interface{} {
	panic("Not implemented")
}

// TODO
func (r *RestFn) ApplyTo(args ISeq) interface{} {
	return nil
}

//TODO
func (r *RestFn) Invoke(args ...interface{}) interface{} {
	return nil
}

func (r *RestFn) ontoArrayPrepend(array []interface{}, args ...interface{}) ISeq {
	return nil
}

func (r *RestFn) findKey(key interface{}, args ISeq) ISeq {
	return nil
}
