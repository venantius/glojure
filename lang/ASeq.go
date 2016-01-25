package lang

// NOTE: Implements ISeq, Sequential, List, Serializable, IHashEq
type ASeq struct {
	Obj

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

// TODO
func (s *ASeq) Equals(ob interface{}) bool {
	return true
}

//TODO
func (s *ASeq) HashCode() int {
	return 0
}

// TODO
func (s *ASeq) HashEq() int {
	return 0
}

// TODO
func (s *ASeq) Count() int {
	return 0
}

func (s *ASeq) Seq() ISeq {
	return s
}

func (s *ASeq) Cons(i interface{}) IPersistentCollection {
	return &Cons{_first: i, _more: s}
}

// TODO: Implement
func (s *ASeq) More() ISeq {
	return nil
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
