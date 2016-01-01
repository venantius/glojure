package lang

// NOTE: Implements Serializable
type Cons struct {
	*ASeq

	_first interface{}
	_more  ISeq
}
