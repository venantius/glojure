package lang

import (
	"reflect"
	"sync"
	"fmt"

)

/*
	Namespace

	Implements: Serializable
*/

// TODO: These may require an "init()" function
var namespacesLock sync.RWMutex      // Protect `namespaces` with a lock
var namespaces map[*Symbol]*Namespace = make(map[*Symbol]*Namespace)// NOTE: This is a ConcurrentHashMap in Java

type Namespace struct {
	AReference

	_meta    IPersistentMap
	name     *Symbol
	mappings IPersistentMap // TODO: AtomicReference
	aliases  IPersistentMap // TODO: AtomicReference

}

func (n *Namespace) String() string {
	return n.name.String()
}

func (n *Namespace) All() ISeq {
	return RT.Seq(MapKeys(namespaces))
}

func (n *Namespace) GetName() *Symbol {
	return n.name
}

func (n *Namespace) GetMappings() IPersistentMap {
	// TODO: AtomicReference, etc.
	return n.mappings
}

func (n *Namespace) Intern(sym *Symbol) *Var {
	if sym.ns != "" {
		panic("Can't intern namespace-qualified symbol")
	}
	m := n.GetMappings()
	var o interface{}
	var v *Var
	for ; o == nil; o = m.ValAt(sym, nil) {
		if v == nil {
			v = &Var{
				ns: n,
				sym: sym,
			}
		}
		// In essence, re-set the current namespace with the new symbol, atomically.
		var newMap IPersistentMap = m.Assoc(sym, v).(IPersistentMap)
		// TODO: n.mappings.compareandset(map, newMap)
		n.mappings = newMap
		m = n.GetMappings()
	}

	switch obj := o.(type) {
	case *Var:
		if obj.ns == n {
			return obj
		}
	}
	if v == nil {
		v = &Var{
			ns: n,
			sym: sym,
		}
	}

	n.warnOrFailOnReplace(sym, o, v)

	// TODO: atomic compareAndSet while loop here
	m = n.GetMappings()

	return v
}

func (n *Namespace) warnOrFailOnReplace(sym *Symbol, o interface{}, v interface{}) {
	switch obj := o.(type) {
	case Var:
		ns := obj.ns
		if ns == n {
			return
		}
		switch vobj := v.(type) {
		case Var:
			if vobj.ns == CLOJURE_NS {
				return
			}
		}
		if ns != CLOJURE_NS {
			panic(fmt.Sprintf("WARNING: %v already refers to: %v in namespace: %v", sym, o, n.name))
		}
	}
	// TODO...the rest of this
}

// TODO
func (n *Namespace) Reference(sym *Symbol, val interface{}) interface{} {
	return nil
}

// TODO
func areDifferentInstancesOfSameType(a interface{}, b interface{}) bool {
	return false
}

// TODO
func referenceType(sym *Symbol, val interface{}) reflect.Type {
	return nil
}

// TODO
func (n *Namespace) Unmap(sym *Symbol) {

}

// TODO: also overloaded
func importType() reflect.Type {
	return nil
}

// TODO
func (n *Namespace) Refer(sym *Symbol, v *Var) *Var {
	return n.Reference(sym, v).(*Var)
}

// TODO
func FindOrCreateNamespace(name *Symbol) *Namespace {
	// TODO: Uncomment all locks
	// namespacesLock.RLock()
	// defer namespacesLock.RUnlock()

	ns := namespaces[name]

	if ns != nil {
		return ns
	}
	newns := &Namespace{
		_meta: nil,
		name:  name,
		// both of these are set atomically
		mappings: DEFAULT_IMPORTS,
		aliases:  RT.Map(),
	}
	ns = putNewNamespace(name, newns)
	if ns == nil {
		return newns
	}
	return ns
}

// Safely set a new namespace. If the key already exists, don't overwrite it.
// The global lock should already be in place when this function is called.
func putNewNamespace(k *Symbol, v *Namespace) *Namespace {
	oldv := namespaces[k]
	if oldv == nil {
		namespaces[k] = v
		return nil
	}
	return oldv
}

// TODO
func RemoveNamespace(name *Symbol) *Namespace {
	return nil
}

func FindNamespace(name *Symbol) *Namespace {
	// namespacesLock.RLock()
	ns := namespaces[name]
	// namespacesLock.RUnlock()
	return ns
}

// TODO
func (n *Namespace) GetMapping(name *Symbol) interface{} {
	return nil
}

// TODO
func (n *Namespace) FindInternedVar(symbol *Symbol) *Var {
	return nil
}

// TODO
func (n *Namespace) GetAliases() IPersistentMap {
	return nil
}

// TODO
func (n *Namespace) LookupAlias(alias *Symbol) *Namespace {
	return nil
}

// TODO
func (n *Namespace) AddAlias(alias *Symbol, ns *Namespace) {

}

// TODO
func (n *Namespace) RemoveAlias(alias *Symbol) {

}

// TODO
func (n *Namespace) readResolve() {

}
