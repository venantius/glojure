package lang

// TODO: leaving this as a major TODO for now....

type IFn interface {
	Invoke(args ...interface{}) interface{}
	// TODO A lot of additional functions are left un-implemented
}
