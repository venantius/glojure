package lang

type IChunkedSeq interface {
	Sequential
	ISeq

	ChunkedFirst() IChunk
	ChunkedNext() ISeq
	ChunkedMore() ISeq
}
