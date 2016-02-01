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
	// Start with a PersistentArrayMap
	r := strings.NewReader("{:a :b :c :d}")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	a := make([]interface{}, 4)
	a[0] = lang.InternKeywordByNsName("a")
	a[1] = lang.InternKeywordByNsName("b")
	a[2] = lang.InternKeywordByNsName("c")
	a[3] = lang.InternKeywordByNsName("d")
	m := lang.CreatePersistentArrayMapWithCheck(a)

	if !(m.Equals(y)) {
		t.Error("Failed to initialize array maps that should have been equal.")
	}

	// Let's do a PersistentHashMap next
	r2 := strings.NewReader("{:a :b :c :d :e :f :g :h :i :j :k :l :m :n :o :p :q :r :s :t :u :v :w :x :y :z}")
	y2 := lang.CreateLispReader(r2).Read(false, io.EOF, rune(0), nil, false, nil, nil)
	var phm []interface{}

	phm = []interface{}{
		lang.InternKeywordByNsName("a"),
		lang.InternKeywordByNsName("b"),
		lang.InternKeywordByNsName("a"), // Note that we're implicitly checking that duplicate
		lang.InternKeywordByNsName("b"), // keys and values don't do weird things.
		lang.InternKeywordByNsName("c"),
		lang.InternKeywordByNsName("d"),
		lang.InternKeywordByNsName("e"),
		lang.InternKeywordByNsName("f"),
		lang.InternKeywordByNsName("g"),
		lang.InternKeywordByNsName("h"),
		lang.InternKeywordByNsName("i"),
		lang.InternKeywordByNsName("j"),
		lang.InternKeywordByNsName("k"),
		lang.InternKeywordByNsName("l"),
		lang.InternKeywordByNsName("m"),
		lang.InternKeywordByNsName("n"),
		lang.InternKeywordByNsName("o"),
		lang.InternKeywordByNsName("p"),
		lang.InternKeywordByNsName("q"),
		lang.InternKeywordByNsName("r"),
		lang.InternKeywordByNsName("s"),
		lang.InternKeywordByNsName("t"),
		lang.InternKeywordByNsName("u"),
		lang.InternKeywordByNsName("v"),
		lang.InternKeywordByNsName("w"),
		lang.InternKeywordByNsName("x"),
		lang.InternKeywordByNsName("y"),
		lang.InternKeywordByNsName("z"),
	}

	phm2 := lang.CreatePersistentHashMap(phm...)

	if !phm2.Equals(y2) {
		t.Error("Failed to initailize hash maps that should have been equal")
	}
}

// Test basic persistent lists

func TestListReader(t *testing.T) {
	// Start with a PersistentArrayMap
	r := strings.NewReader("(+ 1 1)")
	y := lang.CreateLispReader(r).Read(false, io.EOF, rune(0), nil, false, nil, nil)

	fmt.Println("Y2",
		// y,
		reflect.TypeOf(y))

	fmt.Printf("%#v", y.(*lang.PersistentList).First())

	if !y.(*lang.PersistentList).Equals(true) {
		t.Error("Failed to read persistent list")
	}
}