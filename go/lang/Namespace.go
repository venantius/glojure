package lang

import (
	"reflect"
	"sync"
)

/*
	Namespace

	Implements: Serializable
*/

// TODO: These may require an "init()" function
var namespacesLock *sync.RWMutex      // Protect `namespaces` with a lock
var namespaces map[*Symbol]*Namespace // NOTE: This is a ConcurrentHashMap in Java

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

// TODO
func (n *Namespace) GetMappings() IPersistentMap {
	return nil
}

// TODO
func (n *Namespace) Intern(sym *Symbol) *Var {
	return nil
}

// TODO
func (n *Namespace) warnOrFailOnReplace(sym *Symbol, o interface{}, v interface{}) {
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
	namespacesLock.RLock()
	ns := namespaces[name]
	namespacesLock.RUnlock()
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
	ns = safelyPutNewNamespace(name, newns)
	if ns == nil {
		return newns
	}
	return ns
}

// Safely set a new namespace. If the key already exists, don't overwrite it.
func safelyPutNewNamespace(k *Symbol, v *Namespace) *Namespace {
	namespacesLock.RLock()
	namespacesLock.Lock()
	oldv := namespaces[k]
	if oldv == nil {
		namespaces[k] = v
		namespacesLock.RUnlock()
		namespacesLock.Unlock()
		return nil
	}
	namespacesLock.RUnlock()
	namespacesLock.Unlock()
	return oldv
}

// TODO
func RemoveNamespace(name *Symbol) *Namespace {
	return nil
}

func FindNamespace(name *Symbol) *Namespace {
	namespacesLock.RLock()
	ns := namespaces[name]
	namespacesLock.RUnlock()
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
