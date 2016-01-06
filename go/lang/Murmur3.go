package lang

import (
	"bytes"
	"encoding/binary"
	"murmur3"
)

type mmh3 struct{}

var Murmur3 = mmh3{}

// Hash an int32
func HashInt(input int32) int {
	if input == 0 {
		return 0
	}
	buf := new(bytes.Buffer)
	binary.Write(
		buf,
		binary.LittleEndian,
		input,
	)
	return hashBytes(buf.Bytes())
}

func HashLong(input int64) int {
	if input == 0 {
		return 0
	}
	buf := new(bytes.Buffer)
	binary.Write(
		buf,
		binary.LittleEndian,
		input,
	)
	return hashBytes(buf.Bytes())
}

func HashString(input string) int {
	return hashBytes([]byte(input))
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

func hashBytes(input []byte) int {
	x := murmur3.New32()
	x.Write(input)

	var out int32
	binary.Read(
		bytes.NewReader(x.Sum(nil)),
		binary.LittleEndian,
		&out,
	)
	return int(out)
}
