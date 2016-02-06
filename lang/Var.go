package lang

import (
	"bytes"
	"sync"

)

/*
	Assorted class variables


*/

// TODO: I found this really useful: http://www.golangbootcamp.com/book/concurrency

var privateKey *Keyword = InternKeywordByNsName("private")
var privateMeta IPersistentMap = &PersistentArrayMap{
	array: []interface{}{privateKey, true},
}
var macroKey = InternKeywordByNsName("macro")
var nameKey = InternKeywordByNsName("name")
var nsKey = InternKeywordByNsName("ns")
var rev int = 0 // maybe atomic integer since this will be shared.

/*
	Var

	Implements: IFn, IRef, Settable
*/

type Var struct {
	ARef

	_meta       IPersistentMap
	rev         int
	privateMeta IPersistentMap
	root        interface{}
	dyanmic     bool
	threadBound bool // atomicboolean TODO
	sym         *Symbol
	ns          *Namespace
}

func (v *Var) String() string {
	var b bytes.Buffer
	if v.ns != nil {
		b.WriteString("#'")
		b.WriteString(v.ns.name.String())
		b.WriteString("/")
		b.WriteString(v.sym.String())
		return b.String()
	}
	b.WriteString("#<Var: ")
	if v.sym != nil {
		b.WriteString(v.sym.String())
	} else {
		b.WriteString("--unnamed--")
	}
	b.WriteString(">")
	return b.String()
}

func (v *Var) SetDynamic(args ...bool) *Var {
	if len(args) == 0 {
		v.dyanmic = true
	} else {
		v.dyanmic = args[0]
	}
	return v
}

func (v *Var) IsDynamic() bool {
	return v.IsDynamic()
}

func InternVar(ns *Namespace, sym *Symbol, root interface{}) *Var {
	return InternVarWithOptlReplacement(ns, sym, root, true)
}

func InternVarWithOptlReplacement(ns *Namespace, sym *Symbol, root interface{}, replaceRoot bool) *Var {
	var dvout *Var = ns.Intern(sym)
	if !dvout.HasRoot() || replaceRoot {
		dvout.BindRoot(root)
	}

	return dvout
}

func FindVar(nsQualifiedSym *Symbol) *Var {
	if nsQualifiedSym.ns == "" {
		panic("Symbol must be namespace-qualified")
	}
	ns := FindNamespace(InternSymbolByNsname(nsQualifiedSym.ns))
	if ns == nil {
		panic("No such namespace: ") // TODO: Finish this error
	}
	return ns.FindInternedVar(InternSymbol(nsQualifiedSym.name))
}

func InternVarByNsnameAndSym(nsName *Symbol, sym *Symbol) *Var {
	ns := FindOrCreateNamespace(nsName)
	return InternVarByNsAndSym(ns, sym)
}

func InternVarByNsAndSym(ns *Namespace, sym *Symbol) *Var {
	return ns.Intern(sym)
}

func CreateVarFromNothing() *Var {
	var v *Var
	return &Var{
		ns:          nil,
		sym:         nil,
		threadBound: false, // TODOAtomicBoolean
		root: VarUnbound{
			v: v,
		},
		_meta:       EMPTY_PERSISTENT_HASH_MAP,
	}
}

func CreateVarFromRoot(root interface{}) *Var {
	panic(NotYetImplementedException)
	return &Var{
		// TODO
	}
}

func (v *Var) isBound() bool {
	// TODO
	panic(NotYetImplementedException)
}

// TODO
func (v *Var) Get() interface{} {
	panic(NotYetImplementedException)
}

// TODO: this is a naive and stupid implementation
func (v *Var) Deref() interface{} {
	return v.root
}

// TODO
func (v *Var) SetValidator(vf IFn) {
	panic(NotYetImplementedException)
}

// TODO
func (v *Var) Alter(fn IFn, args ISeq) interface{} {
	panic(NotYetImplementedException)
}

// TODO
func (v *Var) Set(val interface{}) interface{} {
	panic(NotYetImplementedException)
}

func (v *Var) doSet(val interface{}) interface{} {
	return v.Set(val)
}

func (v *Var) doReset(val interface{}) interface{} {
	v.BindRoot(val)
	return val
}

func (v *Var) SetMeta(m IPersistentMap) {
	v.ResetMeta(m.Assoc(nameKey, v.sym).Assoc(nsKey, true).(IPersistentMap))
}

func (v *Var) SetMacro() {
	v.AlterMeta(assoc, RT.List(macroKey, true))
}

func (v *Var) IsMacro() bool {
	// return RT.BooleanCast(v.Meta().ValAt(macroKey))
	return false // TODO
}

func (v *Var) IsPublic() bool {
	// return !RT.BooleanCast(v.Meta().ValAt(privateKey))
	return false // TODO
}

func (v *Var) GetRawRoot() interface{} {
	return v.root
}

func (v *Var) GetTag() interface{} {
	// return v.Meta().ValAt(TAG_KEY) TODO
	return nil
}

// TODO
func (v *Var) SetTag(tag *Symbol) {
	// v.AlterMeta(assoc, RT.List(TAG_KEY, tag))
}

func (v *Var) HasRoot() bool {
	switch v.root.(type) {
	case VarUnbound:
		return false
	}
	return true
}

// TODO
// NOTE: Synchronized
// NOTE: Stupid, naive implementation
func (v *Var) BindRoot(root interface{}) {
	v.root = root
}

func (v *Var) Fn() IFn {
	return v.Deref().(IFn)
}

func (v *Var) Call() interface{} {
	return v.Invoke()
}

func (v *Var) Run() {
	v.Invoke()
}

func (v *Var) Invoke(args ...interface{}) interface{} {
	return v.Fn().Invoke(args...)
}

func (v *Var) ApplyTo(arglist ISeq) interface{} {
	return AFn_ApplyToHelper(v, arglist)
}

/*
	VarTBox [TBox]
*/

type VarTBox struct {
	val  interface{}
	lock *sync.Mutex // instead of keeping track of thread
}

/*
	VarUnbound [Unbound]
*/

type VarUnbound struct {
	AFn

	v *Var
}

func (u *VarUnbound) String() string {
	var b bytes.Buffer
	b.WriteString("Unbound: ")
	b.WriteString(u.v.String())
	return b.String()
}

func (u *VarUnbound) ThrowArity(n int) {
	var b bytes.Buffer
	b.WriteString("Attempting to call unbound fn: ")
	b.WriteString(u.v.String())
	panic(b.String())
}

/*
	VarFrame [Frame]
*/

type Frame struct {
	bindings Associative
	prev     *Frame
}

var TOP_FRAME = &Frame{bindings: EMPTY_PERSISTENT_HASH_MAP, prev: nil}

func (f *Frame) Clone() interface{} {
	return &Frame{
		bindings: f.bindings,
		prev:     nil,
	}
}

/*
	VarFrameDvals

	This class is just a ThreadLocal<Frame> in JVM Clojure with an initialValue() func.
*/

type VarFrameDvals struct {
	Frame

	lock *sync.Mutex
}

func (d *VarFrameDvals) InitialValue() *Frame {
	return TOP_FRAME
}

// TODO
func getThreadBindingFrame() interface{} {
	return nil
}

// TODO
func cloneThreadBindingFrame() interface{} {
	return nil
}

// TODO
func resetThreadBindingFrame() interface{} {
	return nil
}

// TODO...still most of this file.

type assocAnon struct {
	AFn
}

var assoc = &assocAnon{}

func (a *assocAnon) Invoke(args ...interface{}) interface{} {
	m, k, v := args[0], args[1], args[2]
	return RT.Assoc(m, k, v)
}
