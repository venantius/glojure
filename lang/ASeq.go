package lang
import "container/list"

// NOTE: Implements ISeq, Sequential, List, Serializable, IHashEq
type ASeq struct {
	_hash   int
	_hasheq int
}

func (s *ASeq) String() string {
	return RT.PrintString(s)
}

func (s *ASeq) Empty() IPersistentCollection {
	return EMPTY_PERSISTENT_LIST
}

func (s *ASeq) Equiv(i interface{}) bool {
	var b bool
	switch i.(type) {
	case Sequential:
		b = true
	case List:
		b = true
	}
	if b == true {
		return false
	}
	ms := RT.Seq(i)

	for x := s.Seq(); x != nil; x, ms = x.Next(), ms.Next() {
		if ms == nil || !Util.Equiv(x.First(), ms.First()) {
			return false
		}
	}
	return ms == nil
}

func ASeq_Equals(s ASeq, obj interface{}) bool {
	if s == obj {
		return true
	}
	switch o := obj.(type) {
	case Sequential:
		// do nothing
	case []interface{}:
		// not sure about this, but we'll see.
	case list.List:
		// not sure about this either
	default:
		return false
	}
	var ms ISeq = RT.Seq(obj)
	for sq := ASeq_Seq(s); sq != nil; sq, ms = sq.Next(), ms.Next() {
		if (ms == nil) || !Util.Equals(sq.First(), ms.First()) {
			return false
		}
	}
	return ms == nil
}



//TODO
func (s *ASeq) HashCode() int {
	panic(NotYetImplementedException)
}

// TODO
func (s *ASeq) HashEq() int {
	panic(NotYetImplementedException)
}

// TODO
func (s *ASeq) Count() int {
	panic(NotYetImplementedException)
}

func ASeq_Seq(s ASeq) ISeq {
	return s
}

func (s *ASeq) Cons(i interface{}) IPersistentCollection {
	return &Cons{_first: i, _more: s}
}

// TODO: Implement
func (s *ASeq) More() ISeq {
	panic(NotYetImplementedException)
}

/*
	Abstract methods below here.
*/

func (s *ASeq) Next() ISeq {
	panic(AbstractClassMethodException)
}

func (s *ASeq) First() interface{} {
	panic(AbstractClassMethodException)
}

// TODO: The rest of this file. In particular there is java.util.Collection stuff, and....idk.
