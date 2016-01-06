package lang

import (
	"fmt"
	"testing"
)

// TODO
func TestAtom(t *testing.T) {
	a := Atom{}
	a.Reset(5)
	fmt.Println(a.Deref())
}
