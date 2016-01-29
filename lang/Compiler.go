package lang

type compiler struct{}

var Compiler = &compiler{}

func (_ *compiler) CurrentNS() *Namespace {
	return CURRENT_NS.Deref().(*Namespace)
}

func (_ *compiler) NamespaceFor(inns *Namespace, sym *Symbol) *Namespace {
	nsSym := InternSymbolByNsname(sym.ns)
	ns := inns.LookupAlias(nsSym)
	if ns == nil {
		ns = FindNamespace(nsSym)
	}
	return ns
}
