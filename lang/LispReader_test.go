package lang_test

import (
	"strings"
	"testing"
	"glojure/lang"
	"io"
	"fmt"
	"reflect"
)

/*
	Vector initialization
 */

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
	r := strings.NewReader("[:asdf ::clojure.core/bool :shenanigans]")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	varArgsVector := lang.CreateVector(
		lang.InternKeywordByNsName("asdf"),
		lang.InternKeywordByNsAndName("clojure.core", "bool"),
		lang.InternKeywordByNsName("shenanigans"),
	)
	if !(varArgsVector.Equals(y)) {
		t.Error("Failed to initialize vectors that should have been equal.")
	}
}

/*
	Map initialization
 */

func TestMapReaderWithKeyword(t *testing.T) {
	r := strings.NewReader("{:a :b :c :d}")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	a := make([]interface{}, 4)
	a[0] = lang.InternKeywordByNsName("a")
	a[1] = lang.InternKeywordByNsName("b")
	a[2] = lang.InternKeywordByNsName("c")
	a[3] = lang.InternKeywordByNsName("d")
	m := lang.CreatePersistentArrayMapWithCheck(a)
	fmt.Println(y, reflect.TypeOf(y))

	fmt.Println(m, reflect.TypeOf(m))
	if !(m.Equals(y)) {
		t.Error("Failed to initialize array maps that should have been equal.")
	}
}