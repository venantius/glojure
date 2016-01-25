package lang

import (
	"testing"
	"fmt"
)

func TestCreateVector(t *testing.T) {
	varArgsVector := CreateVector(1, 2, "asdf")

	items := make([]interface{}, 3)
	items[0] = 1
	items[1] = 2
	items[2] = "asdf"
	interfaceSliceVector := CreateVector(items)

	fmt.Println(interfaceSliceVector, varArgsVector)


	if !(varArgsVector.Equals(interfaceSliceVector)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}
}

func TestCons(t *testing.T) {
	initial := CreateVector(1, 2)

	result := initial.Cons(3)
	if result.Count() != 3 {
		t.Error("Failed to add element to PersistentVector with Cons")
	}
}
