package lang

type AFn struct{}

func (a *AFn) Call() interface{} {
	return a.Invoke()
}

func (a *AFn) Run() {
	a.Invoke()
}

// This should be overwritten by the sub-type
// TODO: some level of implementation detail
func (a *AFn) Invoke(args ...interface{}) interface{} {
	panic(AbstractClassMethodException)
}

// TODO: Implement me!
func (a *AFn) ApplyTo(arglist ISeq) interface{} {
	return nil
}

// TODO: Implement me!
func ApplyToHelper(ifn IFn, arglist ISeq) interface{} {
	return nil
}

// TODO: Implement me!
func (a *AFn) ThrowArity(n int) interface{} {
	return nil
}
