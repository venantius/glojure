package lang
import (
	"container/list"

)

// NOTE: Implements ISeq, Sequential, List, Serializable, IHashEq
type ASeq interface {
	ISeq

	Equals(i interface{}) bool
	HashCode() int
	HashEq() int
	String() string

}

func ASeq_String(s ASeq) string {
	return RT.PrintString(s)
}

func ASeq_Empty(s ASeq) IPersistentCollection {
	return EMPTY_PERSISTENT_LIST
}

func ASeq_Equiv(s ASeq, i interface{}) bool {
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
		if o == 1 {
			// do nothing
		}
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
func ASeq_HashCode(s ASeq) int {
	panic(NotYetImplementedException)
}

// TODO
func ASeq_HashEq(s ASeq) int {
	panic(NotYetImplementedException)
}

// TODO
func ASeq_Count(s ASeq) int {
	panic(NotYetImplementedException)
}

func ASeq_Seq(s ASeq) ISeq {
	return s.(ISeq)
}

func ASeq_Cons(s ASeq, i interface{}) IPersistentCollection {
	return &Cons{_first: i, _more: s}
}

// TODO: Implement
func ASeq_More(s ASeq) ISeq {
	panic(NotYetImplementedException)
}


// TODO: The rest of this file. In particular there is java.util.Collection stuff, and....idk.
