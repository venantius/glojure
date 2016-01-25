package lang_test

import (
	"strings"
	"testing"
	"glojure/lang"
	"io"
	"fmt"
	"reflect"
)

func TestVectorReaderWithInts(t *testing.T) {
	r := strings.NewReader("[1 2 5]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := lang.CreateVector(1, 2, 5)
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}
}

func TestVectorReaderWithStrings(t *testing.T) {
	r := strings.NewReader("[\"a\" \"b\" \"c\"]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := lang.CreateVector("a", "b", "c")
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}
}

func TestVectorReaderWithKeyword(t *testing.T) {
	r := strings.NewReader("[:asdf ::asdf/bool :shenanigans]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	fmt.Println(y, reflect.TypeOf(y))
	varArgsVector := lang.CreateVector("a", "b", "c")
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}
}