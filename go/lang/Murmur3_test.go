package lang

import (
	"testing"
)

func TestHashInt(t *testing.T) {
	if HashInt(1) != 708899323 {
		t.Error("HashInt failed")
	}
}

func TestHashLong(t *testing.T) {
	if HashLong(708899323234234234) != 840551126 {
		t.Error("HashLong failed.")
	}
}

func TestHashString(t *testing.T) {
	if HashString("asdf") != 652222491 {
		t.Error("HashString failed.")
	}
}

// TODO: Add more tests.
