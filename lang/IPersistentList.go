package lang

type IPersistentList interface {
	Sequential
	IPersistentStack
	WithMeta(meta IPersistentMap) IPersistentList
}
