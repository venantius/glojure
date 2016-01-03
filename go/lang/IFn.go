package lang

// TODO: leaving this as a major TODO for now....

type IFn interface {
	Invoke(args ...interface{}) interface{}
	ApplyTo(arglist ISeq) interface{}

	// TODO A lot of additional functions are left un-implemented, but I don't understand them
}
