package lang

// In case it wasn't obvious (it wasn't to me), RT stands for RunTime
type RT struct{}

func IsReduced(r interface{}) bool {
	switch r.(type) {
	case Reduced:
		return true
	default:
		return false
	}
}
