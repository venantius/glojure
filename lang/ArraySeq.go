package lang

/*
	ArraySeq

	Extends abstract class ASeq

	Implements: IndexedSeq, IReduce
 */

type ArraySeq struct {
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

/*
	Abstract methods (ArraySeq)
 */

func (s *ArraySeq) Cons(o interface{}) IPersistentCollection {
	return ASeq_Cons(s, o)
}

func (s *ArraySeq) Empty() IPersistentCollection {
	return ASeq_Empty(s)
}

func (s *ArraySeq) Equals(o interface{}) bool {
	return ASeq_Equals(s, o)
}

func (s *ArraySeq) Equiv(o interface{}) bool {
	return ASeq_Equiv(s, o)
}

func (s *ArraySeq) HashCode() int {
	return ASeq_HashCode(s)
}

func (s *ArraySeq) HashEq() int {
	return ASeq_HashEq(s)
}

func (s *ArraySeq) More() ISeq {
	return ASeq_More(s)
}

func (s *ArraySeq) Seq() ISeq {
	return ASeq_Seq(s)
}

func (s *ArraySeq) String() string {
	return ASeq_String(s)
}

// TODO: There is a fuck ton more stuff in this class that I haven't implmemented.
