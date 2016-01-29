package lang

type IProxy interface {
	__InitClojureFnMappings(m IPersistentMap)
	__UpdateClojureFnMappings(m IPersistentMap)
	__GetClojureFnMappings() IPersistentMap
}
