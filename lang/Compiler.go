package lang

import (
	"io"
	"math/rand"
	"golang.org/x/net/context"
)

const (
	STATEMENT = iota
	EXPRESSION = iota
	RETURN = iota
	EVAL = iota
)

/*
	Compiler variables, etc.

	Note that a number of these are defined in LispReader.go. I don't know why they are re-defined here
	in JVM Clojure.
 */

var DEF *Symbol = InternSymbolByNsname("def")
var LOOP *Symbol = InternSymbolByNsname("loop*")
var RECUR *Symbol = InternSymbolByNsname("recur")
var IF *Symbol = InternSymbolByNsname("if")
var LET *Symbol = InternSymbolByNsname("let*")
var LETFN *Symbol = InternSymbolByNsname("letfn*")
var DO *Symbol = InternSymbolByNsname("do")
var FN *Symbol = InternSymbolByNsname("fn*")
var FNONCE *Symbol = InternSymbolByNsname("fn*").WithMeta(RT.Map(InternKeywordByNsName("once"), true)).(*Symbol)
var DOT *Symbol = InternSymbolByNsname(".")
var ASSIGN *Symbol = InternSymbolByNsname("set!")
var TRY *Symbol = InternSymbolByNsname("try")
var CATCH *Symbol = InternSymbolByNsname("catch")
var FINALLY *Symbol = InternSymbolByNsname("finally")
var THROW *Symbol = InternSymbolByNsname("throw")
var MONITOR_ENTER *Symbol = InternSymbolByNsname("monitor-enter")
var MONITOR_EXIT *Symbol = InternSymbolByNsname("monitor-exit")
var IMPORT *Symbol = InternSymbolByNsAndName("clojure.core", "import")
var DEFTYPE *Symbol = InternSymbolByNsname("deftype*")
var CASE *Symbol = InternSymbolByNsname("case*")
var CLASS *Symbol = InternSymbolByNsname("Class") // TODO I feel like this might end up being irrelevant
var NEW *Symbol = InternSymbolByNsname("new")
var THIS *Symbol = InternSymbolByNsname("this") // TODO unused?
var REIFY *Symbol = InternSymbolByNsname("reify*")
var IDENTITY *Symbol = InternSymbolByNsAndName("clojure.core", "identity")
var _AMP_ *Symbol = InternSymbolByNsname("&")
var ISEQ *Symbol = InternSymbolByNsname("clojure.lang.ISeq")
var loadNs *Symbol = InternSymbolByNsname("load-ns")
var inlineKey *Symbol = InternSymbolByNsname("inline")
// TODO: more declarations here...

var LOCAL_ENV *Var = CreateVarFromRoot(nil).SetDynamic()
var LOOP_LOCALS *Var = CreateVarFromNothing().SetDynamic()
var LOOP_LABEL *Var = CreateVarFromNothing().SetDynamic()
var CONSTANTS *Var = CreateVarFromNothing().SetDynamic()
var CONSTANT_IDS *Var = CreateVarFromNothing().SetDynamic()
var KEYWORD_CALLSITES *Var = CreateVarFromNothing().SetDynamic()
var PROTOCOL_CALLSITES *Var = CreateVarFromNothing().SetDynamic()
var VAR_CALLSITES *Var = CreateVarFromNothing().SetDynamic()
var KEYWORDS *Var = CreateVarFromNothing().SetDynamic()
var VARS *Var = CreateVarFromNothing().SetDynamic()
var METHOD *Var = CreateVarFromRoot(nil).SetDynamic()
var IN_CATCH_FINALLY *Var = CreateVarFromNothing()
// TODO...more declarations.

var COMPILE_PATH *Var = InternVar(FindOrCreateNamespace(InternSymbolByNsname("clojure.core")),
	InternSymbolByNsname("*compile-path*"), nil).SetDynamic()

// TODO..more declarations...
var LINE *Var = CreateVarFromRoot(0).SetDynamic()
var COLUMN *Var = CreateVarFromRoot(0).SetDynamic()

type ObjExpr struct {
	name string
	internalName string
	thisName string
	keywords IPersistentMap
	constants *PersistentVector
}

/*
	Compiler struct and methods
 */

type compiler struct{}

var Compiler = &compiler{}

// TODO
func (_ *compiler) CurrentNS() *Namespace {
	panic(NotYetImplementedException)
	// return CURRENT_NS.Deref().(*Namespace)
}

func (_ *compiler) NamespaceFor(inns *Namespace, sym *Symbol) *Namespace {
	nsSym := InternSymbolByNsname(sym.ns)
	ns := inns.LookupAlias(nsSym)
	if ns == nil {
		ns = FindNamespace(nsSym)
	}
	return ns
}

func (_ *compiler) Macroexpand1(x interface{}) interface{} {
	panic(NotYetImplementedException)
}

func (_ *compiler) Macroexpand(form interface{}) interface{} {
	exf := Compiler.Macroexpand1(form)
	if exf != form {
		return Compiler.Macroexpand(exf)
	}
	return form
}

func (_ *compiler) Compile(rdr *io.Reader, sourcePath string, sourceName string) interface{} {
	// TODO: Do we need this? I don't know.
	// #VESTIGIAL
	/*
	if COMPILE_PATH.Deref() == nil {
		panic("*compile-path* not set")
	}
	*/

	var EOF int = rand.Int() // TODO: Sentinel value
	var ret interface{}

	return 1


	// TODO: Some other stuff in here. Not sure how important it is.

}

// In JVM Clojure, gen is a GeneratorAdapter. We don't have an analog for that here.
func (_ *compiler) Compile1(ctx context.Context, gen interface{}, objx ObjExpr, form interface{}) {
	// TODO: some initial set-up.
	line := LINE.Deref(ctx)
	column := COLUMN.Deref(ctx)

	ctx = pushBindingsForContext(ctx, RT.Map(LINE, line, COLUMN, column)) // ROADMAP: Classloader
	// try block begins here
	form = Compiler.Macroexpand(form)
	switch f := form.(type) {
	case ISeq:
		if Util.Equals(RT.First(form), DO) {
			for s := RT.Next(f); s != nil; s = RT.Next(s) {
				Compiler.Compile1(ctx, gen, objx, RT.First(s))
			}
		}
	default:
		expr := Compiler.Analyze(EVAL, form)
		objx.keywords = KEYWORDS.Deref(ctx).(IPersistentMap)
		objx.vars = VARS.Deref(ctx).(IPersistentMap)
		objx.constants = CONSTANTS.Deref().(*PersistentVector)
		expr.Emit(EXPRESSION, objx, gen)
		expr.Eval()
	}
}

func (_ *compiler) Eval(form interface{}, freshLoader bool) interface{} {
	createdLoader := false // do we need this?
	panic(NotYetImplementedException)
}

func (_ *compiler) Analyze(a interface{}, b interface{}) interface{} {
	panic(NotYetImplementedException)
}