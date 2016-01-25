package lang

// NOTE: Implements IndexedSeq, IReduce
type ArraySeq struct {
	ASeq

	_meta IPersistentMap
	array []interface{}
	i     int
}

// NOTE: This actually returns null in the Java version as well
func CreateArraySeq(array ...interface{}) *ArraySeq {
	if array == nil || len(array) == 0 {
		return nil
	}
	return &ArraySeq{
		array: array,
		i:     0,
	}
}

// TODO
func CreateArraySeqFromObject(array interface{}) ISeq {
	return nil
}

func (a *ArraySeq) First() interface{} {
	if a.array != nil {
		return a.array[a.i]
	}
	return nil
}

func (a *ArraySeq) Next() ISeq {
	if a.array != nil && a.i+1 < len(a.array) {
		return &ArraySeq{
			array: a.array,
			i:     a.i + 1,
		}
	}
	return nil
}

func (a *ArraySeq) Count() int {
	if a.array != nil {
		return len(a.array) - a.i
	}
	return 0
}

func (a *ArraySeq) Index() int {
	return a.i
}

func (a *ArraySeq) WithMeta(meta IPersistentMap) *ArraySeq {
	return &ArraySeq{
		_meta: meta,
		array: a.array,
		i:     a.i,
	}
}

// TODO
func (a *ArraySeq) Reduce(f IFn, start interface{}) interface{} {
	return nil
}

// TODO: There is a fuck ton more stuff in this class that I haven't implmemented.
