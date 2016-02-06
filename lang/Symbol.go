package lang

import (
	"strings"
)

// NOTE: Implements IObj, Comparable, Named, Serializable, IHashEq
type Symbol struct {
	AFn

	ns      string
	name    string
	_hasheq int
	_meta   IPersistentMap
	_str    string
}

func (s *Symbol) String() string {
	if s._str == "" {
		if s.ns != "" {
			s._str = s.ns + "/" + s.name
		} else {
			s._str = s.name
		}
	}
	return s._str
}

func (s *Symbol) GetNamespace() string {
	return s.ns
}

func (s *Symbol) GetName() string {
	return s.name
}

/*
	NOTE: There's a note in the original Java code that `CreateSymbol`
	only exists for backwards compatibility with code compiled against
	earlier versions of Clojure
*/
func CreateSymbol(args ...string) *Symbol {
	return InternSymbol(args...)
}

func InternSymbolByNsAndName(ns string, name string) *Symbol {
	return &Symbol{
		ns:   ns,
		name: name,
	}
}

func InternSymbolByNsname(nsname string) *Symbol {
	i := strings.Index(nsname, "/")
	if i == -1 || nsname == "/" {
		return &Symbol{
			name: nsname,
		}
	} else {
		return &Symbol{
			ns:   nsname[0:i],
			name: nsname[i+1:],
		}
	}
}

func InternSymbol(args ...string) *Symbol {
	if len(args) == 1 {
		return InternSymbolByNsname(args[0])
	} else if len(args) == 2 {
		return InternSymbolByNsAndName(args[0], args[1])
	}
	panic(WrongNumberOfArgumentsException)
}

func (s *Symbol) Equals(obj interface{}) bool {
	if s == obj {
		return true
	}
	switch obj.(type) {
	case *Symbol:
	//continue
	default:
		return false
	}

	symbol := obj.(*Symbol)
	return Util.Equals(s.ns, symbol.ns) && Util.Equals(s.name, symbol.name)
}

func (s *Symbol) HashCode() int {
	return Util.HashCombine(HashString(s.name), Util.Hash(s.ns))
}

func (s *Symbol) HashEq() int {
	if s._hasheq == 0 {
		s._hasheq = Util.HashCombine(HashString(s.name), Util.Hash(s.ns))
	}
	return 0
}

func (s *Symbol) WithMeta(meta IPersistentMap) interface{} {
	return &Symbol{
		_meta: meta,
		ns:    s.ns,
		name:  s.name,
	}
}

func (s *Symbol) CompareTo(obj interface{}) int {
	objS := obj.(Symbol)
	if s.Equals(obj) {
		return 0
	}
	if &s.ns == nil && &objS.ns != nil {
		return -1
	}
	if &s.ns != nil {
		if &objS.ns == nil {
			return 1
		}
		nsc := Util.StringCompareTo(s.ns, objS.ns)
		if nsc != 0 {
			return nsc
		}
	}
	return Util.StringCompareTo(s.name, objS.name)
}

func (s *Symbol) readResolve() interface{} {
	return InternSymbol(s.ns, s.name)
}

func (s *Symbol) Invoke(obj interface{}, notFound interface{}) interface{} {
	return RT.Get(obj, s, notFound)
}

func (s *Symbol) Meta() IPersistentMap {
	return s._meta
}
