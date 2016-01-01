package lang

import (
	"errors"
)

type AFn struct{}

func (a *AFn) call() interface{} {
	return a.invoke()
}

func (a *AFn) run() {
	a.invoke()
}

// This should be overwritten by the sub-type
// TODO: some level of implementation detail
func (a *AFn) invoke(args ...interface{}) interface{} {
	panic(errors.New("Not implemented"))
}

// TODO: Implement me!
func (a *AFn) applyTo(arglist ISeq) interface{} {
	return nil
}

// TODO: Implement me!
func applyToHelper(ifn IFn, arglist ISeq) interface{} {
	return nil
}

// TODO: Implement me!
func (a *AFn) throwArity(n int) interface{} {
	return nil
}
