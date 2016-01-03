package lang

// NOTE: Implements IFn, Comparable, Named, Serializable, IHashEq
type Keyword struct {
	table  int // TODO: Should be ConcurrentHashMap ???
	rq     int // should be ReferenceQueue
	sym    Symbol
	hasheq int
	_str   string
}

/*
	TODO: This method is overloaded in Java
	takes (symbol), (str, str), (str)
*/
func (k *Keyword) Intern(sym Symbol) *Keyword {
	return nil
}

/*
	TODO: This method is ALSO overloaded in Java
	takes (symbol), (str, str), (str)
*/
func (k *Keyword) Find(sym Symbol) *Keyword {
	return nil
}

// TODO
func (k *Keyword) HashCode() int {
	return 0
}

func (k *Keyword) HashEq() int {
	return k.hasheq
}
