package lang

// TODO: Extends Obj
// NOTE: Implements ISeq, Sequential, List, Serializable, IHashEq
type ASeq struct {
	*Obj

	_hash   int
	_hasheq int
}

func (s *ASeq) ConsISeq(i interface{}) ISeq {
	return Cons{_first: i, _more: s}
}

// TODO: Implement
func (s *ASeq) Next() ISeq {
	return nil
}

// TODO: Implement
func (s *ASeq) More() ISeq {
	return nil
}

// TODO: Implement
func (s *ASeq) First() interface{} {
	return nil
}

func (s *ASeq) Equiv(i interface{}) bool {
	var b bool
	switch i.(type) {
	case Sequential:
		b = false
	case List:
		b = false
	}
	if b == false {
		return false
	}
	ms := RT.Seq(i)
	// TODO: some other stuff here
	return ms == nil
}

func (s *ASeq) Seq() ISeq {
	return s
}
