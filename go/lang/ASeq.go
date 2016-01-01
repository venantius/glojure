package lang

// TODO: Extends Obj
// NOTE: Implements ISeq, Sequential, List, Serializable, IHashEq
type ASeq struct {
}

func (s *ASeq) Cons(i interface{}) ISeq {
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
