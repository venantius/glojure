package lang

type IChunkedSeq interface {
	Sequential
	ISeq

	chunkedFirst() IChunk
	chunkedNext() ISeq
	chunkedMore() ISeq
}
