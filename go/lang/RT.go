package lang

import (
	"go/types"
	"os"
)

// In case it wasn't obvious (it wasn't to me), RT stands for RunTime
type rt struct{}

/*
	NOTE: I've made the design decision for now to mimic static methods as best I can,
	which in this case means creating a private class and a single public object for that class. In practice I think that RT.java is more or less a catchall for a host of static methods that could just as easily be generally pure functions.

	I'll decide whether or not I want to change this later.
*/

var T bool = true
var F bool = false
var LOADER_SUFFIX string = "__init"

var DEFAULT_IMPORTS IPersistentMap = RT.Map(
	InternSymbol("Boolean"), types.Bool,
	InternSymbol("Byte"), types.Byte,
	InternSymbol("Rune"), types.Rune, // NOTE "Character" in JVM Clojure
	// TODO ... there's a lot of stuff in here. It basically maps all names to underlying types.
)

func (_ *rt) ReadTrueFalseUnknown(s string) interface{} {
	if s == "true" {
		return true
	} else if s == "false" {
		return false
	}
	return InternKeywordByNsName("unknown")
}

func (_ *rt) GetEnvWithDefault(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	return defaultVal
}

// TODO...more here
var CLOJURE_NS = FindOrCreateNamespace(InternSymbol("clojure.core"))
var readeval interface{} = RT.ReadTrueFalseUnknown(RT.GetEnvWithDefault("clojure.read.eval", "true"))
var READEVAL = InternVar(CLOJURE_NS, InternSymbol("*read-eval*"), readeval).SetDynamic(true)

func (_ *rt) EMPTY_ARRAY() []interface{} {
	return make([]interface{}, 1)
}

var RT = rt{} // Mock static methods

func (_ *rt) Map(init ...interface{}) IPersistentMap {
	if init == nil {
		return EMPTY_PERSISTENT_ARRAY_MAP
	} else if len(init) <= HASHTABLE_THRESHOLD {
		return CreatePersistentArrayMapWithCheck(init)
	}
	return CreatePersistentHashMapWithCheck(init)
}

func (_ *rt) IsReduced(r interface{}) bool {
	switch r.(type) {
	case Reduced:
		return true
	default:
		return false
	}
}

// TODO....so much

func (_ *rt) Seq(coll interface{}) ISeq {
	switch c := coll.(type) {
	case *ASeq:
		return c
	case *LazySeq:
		return c.Seq()
	}
	return RT.seqFrom(coll)
}

func (_ *rt) seqFrom(coll interface{}) ISeq {
	// TODO
	return nil
}

func (_ *rt) SubVec(v IPersistentVector, start int, end int) IPersistentVector {
	if end < start || start < 0 || end > v.Count() {
		panic(IndexOutOfBoundsException)
	}
	if start == end {
		return nil
		// return EMPTY_PERSISTENT_VECTOR TODO
	}
	return &SubVector{} // TODO
}

func (_ *rt) getFrom(coll interface{}, key interface{}, notFound interface{}) interface{} {
	if coll == nil {
		return nil
	}
	// TODO: This implementation is incomplete
	return nil
}

func (_ *rt) Assoc(coll interface{}, key interface{}, val interface{}) Associative {
	if coll == nil {
		array := []interface{}{key, val}
		array[0] = key
		array[1] = val
		return &PersistentArrayMap{array: array}
	}
	return coll.(Associative).Assoc(key, val)
}

func (_ *rt) Get(coll interface{}, key interface{}, notFound interface{}) interface{} {
	switch coll.(type) {
	case ILookup:
		return coll.(ILookup).ValAt(key, notFound)
	}
	return RT.getFrom(coll, key, notFound)
}

// unordered

// TODO
func (_ *rt) Count(o interface{}) int {
	return 0
}

// TODO
func (_ *rt) PrintString(o interface{}) string {
	return ""
}

// TODO
func (_ *rt) ToArray(coll interface{}) []interface{} {
	return nil
}

func (_ *rt) Cons(x interface{}, coll interface{}) ISeq {
	if coll == nil {
		return &PersistentList{
			_first: x,
			_rest:  nil,
			_count: 1,
		}
	}
	switch c := coll.(type) {
	case ISeq:
		return &Cons{
			_first: x,
			_more:  c,
		}
	default:
		return &Cons{
			_first: x,
			_more:  RT.Seq(coll),
		}
	}

}

func (_ *rt) Next(x interface{}) ISeq {
	switch s := x.(type) {
	case ISeq:
		return s.Next()
	}
	seq := RT.Seq(x)
	if seq == nil {
		return nil
	}
	return seq.Next()
}

/*
	List (Persistent) support
*/

func (_ *rt) List(args ...interface{}) ISeq {
	if len(args) == 0 {
		return nil
	} else if len(args) == 1 {
		return &PersistentList{
			_first: args[0],
			_rest:  nil,
			_count: 1,
		}
	} else {
		newarray := make([]interface{}, len(args)+1)
		copy(newarray, args)
		return RT.ListStar(newarray...)
	}
}

func (_ *rt) ListStar(args ...interface{}) ISeq {
	if len(args) == 1 {
		return args[0].(ISeq)
	} else {
		return RT.Cons(args[0], RT.ListStar(args[1:]...))
	}
}
