package lang

import (
	"testing"
)

func TestCreateVector(t *testing.T) {
	varArgsVector := CreateVector(1, 2, "asdf")

	items := make([]interface{}, 3)
	items[0] = 1
	items[1] = 2
	items[2] = "asdf"
	interfaceSliceVector := CreateVector(items)

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

func TestNth(t *testing.T) {
	initial := CreateVector("a", "b")
	result := initial.Nth(0, nil)

	if result != "a" {
		t.Error("Failed to retrieve nth element of PersistentVector")
	}

	notFoundResult := initial.Nth(3, "c")
	if notFoundResult != "c" {
		t.Error("Failed to use notFound value when retrieving nth element of PersistentVector")
	}
}