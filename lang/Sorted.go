package lang

type Sorted interface {
	// TODO Comparator() Comparator
	EntryKey(entry interface{}) interface{}
	Seq(ascending bool) ISeq
	SeqFrom(key interface{}, ascending bool) ISeq
}
