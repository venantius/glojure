package lang_test

import (
	"strings"
	"testing"
	"glojure/go/lang"
	"io"
	"fmt"
	"reflect"
)

func TestVectorReaderWithIntegerVector(t *testing.T) {
	r := strings.NewReader("[1 2 5]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := lang.CreateVector(1, 2, 5)

	//fmt.Println(y, varArgsVector)
	//fmt.Println(reflect.TypeOf(y), reflect.TypeOf(varArgsVector))

	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}


}

func TestVectorReaderWithStringVector(t *testing.T) {
	r := strings.NewReader("[\"a\" \"b\" \"c\"]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	fmt.Println(y, reflect.TypeOf(y))
	varArgsVector := lang.CreateVector("a", "b", "c")
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}


}
