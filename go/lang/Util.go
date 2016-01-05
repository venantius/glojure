package lang

type util struct{}

var Util = util{}

// NOTE / TODO: This is a heavily overloaded function
func (_ *util) Equiv(k1 interface{}, k2 interface{}) bool {
	if k1 == k2 {
		return true
	}
	if k1 != nil {

		// check if TODO...
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
	return k1.Equals(k2)
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
	return k1.(IPersistentCollection).Equals(k2)
}

// Back to Util functions

// TODO
func (_ *util) EquivPred(k1) EquivPred {
	return nil
}

// TODO
func (_ *util) PCEquiv(k1 interface{}, k2 interface{}) bool {
	return false
}

// TODO
func (_ *util) Equals(k1 interface{}, k2 interface{}) bool {
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

/*
	NOTE: Everything after here isn't in the Java version, but I needed it
	and...utils namespace!

	~ @venantius
*/

func (_ *util) StringCompareTo(first string, second string) int {
	// TODO
	return 0
}
