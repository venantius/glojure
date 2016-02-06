package lang

import (
	"bytes"
	"golang.org/x/net/context"
	"fmt"
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

	Extends: ARef

	Implements: IFn, IRef, Settable

	Notes: in JVM Clojure, the implementation of Vars is heavily dependent on the JVM's support
	for thread-local storage. In the almost total absence of a simple analog here in Go, I've
	chosen to instead use the x/net/context package instead as a way of passing explicit
	bindings.
*/

type Var struct {
	ARef

	// AReference
	_meta       IPersistentMap

	// ARef
	validator IFn
	watches *PersistentHashMap // default EMPTY

	rev         int
	privateMeta IPersistentMap
	root        interface{}
	dynamic     bool
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
		v.dynamic = true
	} else {
		v.dynamic = args[0]
	}
	return v
}

func (v *Var) IsDynamic() bool {
	return v.dynamic
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

// unexported type and key for var context utility functions
type key int
var varKey key = 0

// Set the new Var root for the provided context
func setVarBindingForContext(ctx context.Context, val *Frame) context.Context {
	return context.WithValue(ctx, varKey, val)
}

// Retrieve the new Var root from the provided context
func getVarBindingFromContext(ctx context.Context) (*Frame, bool) {
	r, ok := ctx.Value(varKey).(*Frame)
	return r, ok
}

// pushBindingsForContext creates a new *Frame with the provided bindings added to
// the bindings on the provided Context's frame. It returns a new Context with the
// new *Frame.
func pushBindingsForContext(ctx context.Context, bindings Associative) context.Context {
	f, ok := getVarBindingFromContext(ctx)
	if !ok {
		panic("Stored var value wasn't an instance of *lang.Frame")
	}
	var bmap Associative = f.bindings
	for bs := bindings.Seq(); bs != nil; bs = bs.Next() {
		var e IMapEntry = bs.First().(IMapEntry)
		var v *Var = e.Key().(*Var)
		if !v.dynamic {
			panic(fmt.Sprintf("Can't dynamically bind non-dynamic var: %v/%v", v.ns, v.sym))
		}
		v.validate(v.GetValidator(), e.Val())
		bmap = bmap.Assoc(v, e.Val())
	}
	return setVarBindingForContext(ctx, &Frame{
		bindings: bmap,
		prev: f,
	})
}

// popBindingsForContext returns a new Context with the previous *Frame.
func popBindingsForContext(ctx context.Context) context.Context {

}

// IsBound checks to see if the Var has a root or local binding.
func (v *Var) IsBound(ctx context.Context) bool {
	r, ok := getVarBindingFromContext(ctx)
	if ok == false {
		panic("Stored var value wasn't an instance of *lang.Frame")
	}
	return v.HasRoot() || r != nil && r.bindings.ContainsKey(v)
}

// Get retrieves the value stored at the root or local binding of this Var.
func (v *Var) Get(ctx context.Context) interface{} {
	r, ok := getVarBindingFromContext(ctx)
	if ok == false {
		panic("Stored var value wasn't an instance of *lang.Frame")
	}
	if r == nil {
		return v.root
	}
	return v.Deref(ctx)
}

// Deref retreives the value stored at the root or local binding of this Var.
func (v *Var) Deref(ctx context.Context) interface{} {
	r, ok := getVarBindingFromContext(ctx)
	if (r != nil) {
		return r.bindings.EntryAt(v)
	}
	return v.root
}

// TODO
func (v *Var) SetValidator(vf IFn) {
	if v.HasRoot() {
		v.validate()
	}
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

func (v *Var) Fn(ctx context.Context) IFn {
	return v.Deref(ctx).(IFn)
}

func (v *Var) Call(ctx context.Context) interface{} {
	return v.Invoke(ctx)
}

func (v *Var) Run(ctx context.Context) {
	v.Invoke(ctx)
}

func (v *Var) Invoke(ctx context.Context, args ...interface{}) interface{} {
	return v.Fn(ctx).Invoke(args...)
}

func (v *Var) ApplyTo(arglist ISeq) interface{} {
	return AFn_ApplyToHelper(v, arglist)
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

// In JVM Clojure, dvals is a ThreadLocal<Frame>. Here we pass it around in a context.
// First we initialize dvals to TOP_FRAME
var dvals *Frame = TOP_FRAME

// TODO...still most of this file.

type assocAnon struct {
	AFn
}

var assoc = &assocAnon{}

func (a *assocAnon) Invoke(args ...interface{}) interface{} {
	m, k, v := args[0], args[1], args[2]
	return RT.Assoc(m, k, v)
}