package lang

import (
	"strings"
	"testing"

	"io"
)

func TestVectorReaderWithIntegerVector(t *testing.T) {
	r := strings.NewReader("[1 2 5]")
	y := createLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := CreateVector(1, 2, 5)
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}


}

func TestVectorReaderWithStringVector(t *testing.T) {
	r := strings.NewReader("[1]")
	y := createLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := CreateVector("a", "b", "c")
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}


}
