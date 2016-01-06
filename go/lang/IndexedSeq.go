package lang

type IndexedSeq interface {
	ISeq
	Sequential
	// Counted

	Index() int
}
