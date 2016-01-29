package lang

/*
	Keywords involve more concurrent magic than I had previously anticipated;
	the most noticable of which are the combination use of a ReferenceQueue and
	a ConcurrentHashMap where the values are Reference<Keyword>. None of these
	structures have direct analogs in Go -- for the time being, I've punted
	on the issue firmly, but it'll have to be revisited at some point.

	~ @venantius, jan 2016
*/

// NOTE: Implements IFn, Comparable, Named, Serializable, IHashEq
type Keyword struct {
	sym    *Symbol
	hasheq int
	_str   string
}

var rq int                                   // Should be ReferenceQueue
var keywordTable = make(map[Symbol]*Keyword) // Should be ConcurrentHashMap with Reference<Keyword>

// TODO
func InternKeyword(sym *Symbol) *Keyword {
	var k *Keyword
	existingRef := keywordTable[*sym]
	if existingRef == nil {
		// TODO: Util.ClearCache(rq, table)
		if sym.Meta() != nil {
			sym = sym.WithMeta(nil).(*Symbol)
		}
		k = &Keyword{
			sym:    sym,
			hasheq: sym.HashEq() + 0x9e3779b9,
		}

		// TODO
	}
	if existingRef == nil {

		if k != nil {
			return k
		}
		return nil
	}
	//TODO...more
	return nil
}

func InternKeywordByNsAndName(ns string, name string) *Keyword {
	return InternKeyword(InternSymbolByNsAndName(ns, name))
}

func InternKeywordByNsName(nsname string) *Keyword {
	return InternKeyword(InternSymbol(nsname))
}

/*
	TODO: This method is ALSO overloaded in Java
	takes (symbol), (str, str), (str)
*/
func (k *Keyword) Find(sym Symbol) *Keyword {
	return nil
}

func (k *Keyword) HashCode() int {
	return k.sym.HashCode() + 0x9e3779b9
}

func (k *Keyword) HashEq() int {
	return k.hasheq
}

func (k *Keyword) String() string {
	if k._str == "" {
		k._str = ":" + k.sym.String()
	}
	return k._str
}

func (k *Keyword) Equals(o interface{}) bool {
	switch k2 := o.(type) {
	case *Keyword:
		return k.sym.name == k2.sym.name && k.sym.ns == k2.sym.ns
	}
	return false
}

// TODO: A bit longer actually
func (k *Keyword) ThrowArity() interface{} {
	panic(IllegalArgumentException)
}

func (k *Keyword) Call() interface{} {
	return k.ThrowArity()
}

func (k *Keyword) Run() {
	panic(UnsupportedOperationException)
}

func (k *Keyword) Invoke() interface{} {
	return k.ThrowArity()
}

func (k *Keyword) CompareTo(o interface{}) int {
	return k.sym.CompareTo(o.(Keyword).sym)
}

func (k *Keyword) GetNamespace() string {
	return k.sym.GetNamespace()
}

func (k *Keyword) GetName() string {
	return k.sym.GetName()
}

func (k *Keyword) ReadResolve() interface{} {
	return InternKeyword(k.sym)
}

// TODO... the rest of this file?
