package lang

type util struct{}

var Util = util{}

// NOTE / TODO: This is a heavily overloaded function
func (_ *util) Equiv(k1 interface{}, k2 interface{}) bool {
	if k1 == k2 {
		return true
	}
	if k1 != nil {
		if IsNumeric(k1) && IsNumeric(k2) {
			return Numbers.Equal(k1, k2)
		}
		switch t1 := k1.(type) {
		case IPersistentCollection:
			return Util.PCEquiv(k1, k2)
		case IEquals:
			return t1.Equals(k2)
		}
		switch t2 := k2.(type) {
		case IPersistentCollection:
			return Util.PCEquiv(k1, k2)
		case IEquals:
			return t2.Equals(k1)
		}
	}
	return false
}

/*
	Declaration block for a bunch of anonymous classes implementing EquivPred
*/

type EquivPred interface {
	Equiv(k1 interface{}, k2 interface{}) bool
}

type equivNull struct{}

func (_ *equivNull) Equiv(k1 interface{}, k2 interface{}) bool {
	return k2 == nil
}

type equivEquals struct{}

func (_ *equivEquals) Equiv(k1 interface{}, k2 interface{}) bool {
	switch obj := k1.(type) {
	case IEquals:
		return obj.Equals(k2)
	default:
		return obj == k2
	}
}

type equivNumber struct{}

func (_ *equivNumber) Equiv(k1 interface{}, k2 interface{}) bool {
	if IsNumeric(k2) {
		return Numbers.Equal(k1, k2) // TODO: Might need more work here
	}
	return false
}

type equivColl struct{}

func (_ *equivColl) Equiv(k1 interface{}, k2 interface{}) bool {
	coll := false
	switch k1.(type) {
	case IPersistentCollection:
		coll = true
	}
	switch k2.(type) {
	case IPersistentCollection:
		coll = true
	}
	if coll {
		return Util.PCEquiv(k1, k2)
	}
	// return k1.(IPersistentCollection).Equals(k2) TODO
	return false
}

// Back to Util functions

func (_ *util) EquivPred(k1 interface{}) EquivPred {
	if k1 == nil {
		return &equivNull{}
	}
	if IsNumeric(k1) {
		return &equivNumber{}
	}
	switch k1.(type) {
	case string:
		return &equivEquals{}
	case *Symbol:
		return &equivEquals{}
	}
	// TODO: case map or generic collection, return equivColl
	return &equivEquals{}
}

// TODO
func (_ *util) PCEquiv(k1 interface{}, k2 interface{}) bool {
	panic(NotYetImplementedException)

	return false
}

// NOTE: I'm not yet super happy about this, but it'll have to do for now.
func (_ *util) Equals(k1 interface{}, k2 interface{}) bool {
	if k1 == k2 {
		return true
	}
	switch k := k1.(type) {
	case IEquals:
		return k1 != nil && k.Equals(k2)
	}
	switch k := k2.(type) {
	case IEquals:
		return k2 != nil && k.Equals(k1)
	}
	return false
}

// Like Hash* functions, but does nil checking first.
func (_ *util) Hash(i interface{}) int {
	if i == nil {
		return 0
	}
	// TODO: the empty interface doesn't have a HashCode method.
	// We should type-check against some sort of interface that implements
	// HashCode and then otherwise dispatch directly to one of the various
	// hashing methods in Murmur3.go
	return 1 // i.HashCode()
}

// TODO
func (_ *util) HashEq(o interface{}) int {
	if o == nil {
		return 0
	}
 	switch obj := o.(type) {
	case IHashEq:
		return obj.HashEq()
	}
	panic("Tried to call Util.HashEq on something that wasn't hashable.")
	return 0
}

func (_ *util) HashCombine(seed int, hash int) int {
	// JVM Clojure note: "A la boost"
	seed ^= hash + 0x9e3779b9 + (seed << 6) + (seed >> 2)
	return seed
}

func (_ *util) SneakyThrow(t interface{}) {
	if t == nil {
		panic("Null pointer exception") // TODO: lol
	}
	panic(t)
}

/*
	NOTE: Everything after here isn't in the Java version, but I needed it
	and...utils namespace!

	~ @venantius
*/

// TODO
func (_ *util) StringCompareTo(first string, second string) int {
	panic(NotYetImplementedException)
	return 0
}
