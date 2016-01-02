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

/*
TODO: uncomment me
func (p *Primordial) doInvoke(args interface{}) interface{} {
	switch args.(type) {
	case ArraySeq:
		argsarray := args.(ArraySeq).array
		ret := EMPTY_PERSISTENT_LIST
		for i := len(argsarray) - 1; i >= 0; i-- {
			ret = ret.Cons(argsarray[i]).(PersistentList)
		}
		return ret
	}
	// TODO
	return nil
}
*/

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
