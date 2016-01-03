package lang

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

// TODO: method overloading
func CreateSymbol(args ...interface{}) *Symbol {
	return nil
}

// TODO: Method overloading
func Intern(args ...interface{}) *Symbol {
	return nil
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
	return 0
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
