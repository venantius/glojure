package lang

// NOTE: Implements IndexedSeq, IReduce
type ArraySeq struct {
	ASeq

	array []interface{}
	i     int
}

// NOTE: This actually returns null in the Java version as well
func CreateArraySeq() *ArraySeq {
	return nil
}

// TODO: There is a fuck ton more stuff in this class that I haven't implmemented.
