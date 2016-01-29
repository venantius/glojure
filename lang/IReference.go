package lang

type IReference interface {
	IMeta

	AlterMeta(alter IFn, args ISeq) IPersistentMap
	ResetMeta(m IPersistentMap) IPersistentMap
}
