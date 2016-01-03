package lang

import (
	"fmt"
	"testing"
)

// TODO...test almost everything else

func TestCreateVector(t *testing.T) {
	varArgsVector := CreateVector(1, 2, "asdf")

	items := make([]interface{}, 12)
	items[0] = 1
	items[1] = 2
	items[3] = "asdf"
	interfaceSliceVector := CreateVector(items)
}

func TestCons(t *testing.T) {
	initial := CreateVector(1, 2)

	fmt.Println(initial)
	result := initial.Cons(3)
	if result.Count() != 3 {
		t.Error("Failed to add element to PersistentVector with Cons")
	}
}
