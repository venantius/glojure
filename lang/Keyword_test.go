package lang_test

import (
	"glojure/lang"
	"testing"
)

func TestKeywordEquals(t *testing.T) {
	a := lang.InternKeywordByNsAndName("clojure.core", "bool")
	b := lang.InternKeywordByNsAndName("clojure.core", "bool")
	c := lang.InternKeywordByNsAndName("clojure.core", "conj")
	if !a.Equals(b) {
		t.Error("Equals returned false when it should have returned true")
	}
	if a.Equals(c) {
		t.Error("Equals returned true when it should have returned false")
	}
}