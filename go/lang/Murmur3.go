package lang

import (
	"bytes"
	"encoding/binary"
	"murmur3"
)

type mmh3 struct{}

var Murmur3 = mmh3{}

// Murmurhash an int
func HashInt(input int) int {
	if input == 0 {
		return 0
	}
	buf := new(bytes.Buffer)
	binary.Write(
		buf,
		binary.LittleEndian,
		int32(input),
	)

	x := murmur3.New32()
	x.Write(buf.Bytes())

	var out int32
	binary.Read(
		bytes.NewReader(x.Sum(nil)),
		binary.LittleEndian,
		&out,
	)
	return int(out)
}

func HashLong(input int32) int {
	return 0
}

func HashString(input string) int {
	return 0
}

// TODO
func MixCollHash(hash int, count int) int {
	panic(NotYetImplementedException)
}

// TODO
func HashOrdered(xs Iterable) int {
	panic(NotYetImplementedException)
}

// TODO
func HashUnordered(xs Iterable) int {
	panic(NotYetImplementedException)
}
