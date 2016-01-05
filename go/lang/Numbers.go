package lang

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
