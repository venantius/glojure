package lang

import (
	"strings"
)

// NOTE: Implements IObj, Comparable, Named, Serializable, IHashEq
type Symbol struct {
	*AFn

	ns      string
	name    string
	_hasheq int
	_meta   IPersistentMap
	_str    string
}

func (s *Symbol) String() string {
	if &s._str == nil {
		if &s.ns != nil {
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
	return Intern(args...)
}

func InternNsAndName(ns string, name string) *Symbol {
	return &Symbol{
		ns:   ns,
		name: name,
	}
}

func InternNsname(nsname string) *Symbol {
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

func Intern(args ...string) *Symbol {
	if len(args) == 1 {
		return InternNsname(args[0])
	} else if len(args) == 2 {
		return InternNsAndName(args[0], args[1])
	}
	panic(WrongNumberOfArgumentsException)
}

func (s *Symbol) Equals(obj interface{}) bool {
	if s.Equals(obj) {
		return true
	}
	switch obj.(type) {
	case Symbol:
		// continue
	default:
		return false
	}

	// symbol := obj.(Symbol)

	// TODO: Util.equals
	return false
}

// TODO
func (s *Symbol) HashCode() int {
	return 0
}

// TODO
func (s *Symbol) HashEq() int {
	return 0
}

func (s *Symbol) WithMeta(meta IPersistentMap) interface{} {
	return &Symbol{
		_meta: meta,
		ns:    s.ns,
		name:  s.name,
	}
}

// TODO
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
		nsc := StringCompareTo(s.ns, objS.ns)
		if nsc != 0 {
			return nsc
		}
	}
	return StringCompareTo(s.name, objS.name)
}

func (s *Symbol) readResolve() interface{} {
	return Intern(s.ns, s.name)
}

func (s *Symbol) Invoke(obj interface{}, notFound interface{}) interface{} {
	return RT.Get(obj, s, notFound)
}

func (s *Symbol) Meta() IPersistentMap {
	return s._meta
}
