package lang

func Bitcount(i int) int {
	i = i - ((i >> 1) & 0x55555555);
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333);
	i = (i + (i >> 4)) & 0x0f0f0f0f;
	i = i + (i >> 8);
	i = i + (i >> 16);
	return i & 0x3f;
}

func IsInt(o interface{}) bool {
	switch o.(type) {
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case int:
		return true
	default:
		return false
	}
}

func IsUint(o interface{}) bool {
	switch o.(type) {
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case uint:
		return true
	default:
		return false
	}
}

func IsFloat(o interface{}) bool {
	switch o.(type) {
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}

func IsComplex(o interface{}) bool {
	switch o.(type) {
	case complex64:
		return true
	case complex128:
		return true
	default:
		return false
	}
}

func IsNumeric(o interface{}) bool {
	return IsInt(o) || IsUint(o) || IsFloat(o) || IsComplex(o)
}

/*
	Static methods
*/

type num struct{}

var Numbers = num{}

// TODO
func (_ *num) Equal(a interface{}, b interface{}) bool {
	return true
}

// TODO: Everything else in this file
