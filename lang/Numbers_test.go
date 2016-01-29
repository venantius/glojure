package lang

import (
	"testing"
)

func TestIsInt(t *testing.T) {
	var Int int
	var Int8 int8
	var Int16 int16
	var Int32 int32
	var Int64 int64

	Int = 8
	Int8 = 8
	Int16 = 8
	Int32 = 8
	Int64 = 8

	if !(IsInt(Int) || IsInt(Int8) || IsInt(Int16) || IsInt(Int32) || IsInt(Int64)) {
		t.Error("Failed to identify integer")
	}
}

func TestIsUint(t *testing.T) {
	var Uint uint
	var Uint8 uint8
	var Uint16 uint16
	var Uint32 uint32
	var Uint64 uint64

	Uint = 8
	Uint8 = 8
	Uint16 = 8
	Uint32 = 8
	Uint64 = 8

	if !(IsUint(Uint) || IsUint(Uint8) || IsUint(Uint16) || IsUint(Uint32) || IsUint(Uint64)) {
		t.Error("Failed to identify unsigned integer")
	}
}

func TestIsFloat(t *testing.T) {
	var Float32 float32
	var Float64 float64

	Float32 = 8.15798273498273498273498972394
	Float64 = 8.258723948729384729834728937492873498234

	if !(IsFloat(Float32) || IsFloat(Float64)) {
		t.Error("Failed to identify float")
	}
}

func TestIsComplex(t *testing.T) {
	var Complex64 complex64
	var Complex128 complex128

	Complex64 = 8258729348729384728347298374918734
	Complex128 = 8582937492783492873498273498273498273948i

	if !(IsComplex(Complex64) || IsComplex(Complex128)) {
		t.Error("Failed to identify unsigned integer")
	}
}
