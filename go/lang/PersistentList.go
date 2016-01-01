package lang

// NOTE: Implements IPersistentList, IReduce, List, Counted
type PersistentList struct {
	ASeq

	_first interface{}
	_rest  IPersistentList
	_count int
}

type Primordial struct {
	RestFn
}

func (p *Primordial) GetRequiredArity() int {
	return 0
}

var EMPTY_PERSISTENT_LIST = PersistentList{}

// TODO
func (l *PersistentList) Cons(i interface{}) ISeq {
	return &PersistentList{}
}

// TODO
func (l *PersistentList) First() interface{} {
	return 1
}

// TODO
func (l *PersistentList) More() ISeq {
	return &PersistentList{}
}

// TODO
func (l *PersistentList) Next() ISeq {
	return &PersistentList{}
}
