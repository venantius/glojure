package lang

import (
	"bytes"
	"fmt"
)

/*
	Compiler errors
*/

// TODO: go through other abstract classes and replace their panics with this
// exception
var AbstractClassMethodException = "You have tried to call a method on an abstract class that lacks a concrete implementation. This is a compiler error and should be reported."
var NotYetImplementedException = "This function is not yet implemented. This is a compiler error and should be reported."

/*
	Runtime errors
*/

var IllegalArgumentException = "Illegal argument exception."
var IndexOutOfBoundsException = "Index out of bounds."
var UnsupportedOperationException = "Unsupported operation."
var WrongNumberOfArgumentsException = "Wrong number of arguments."

func DuplicateKeyException(s fmt.Stringer) string {
	var buffer bytes.Buffer
	buffer.WriteString("Duplicate key: ")
	buffer.WriteString(s.String())
	return buffer.String()
}

func NoValueSuppliedForKeyException(s fmt.Stringer) string {
	var buffer bytes.Buffer
	buffer.WriteString("No value supplied for key: ")
	buffer.WriteString(s.String())
	return buffer.String()
}
