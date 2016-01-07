package lang

type LazilyPersistentVector struct{}

func CreateOwningLazilyPersistentVector(items ...interface{}) IPersistentVector {
	if len(items) <= NODE_SIZE {
		return &PersistentVector{
			cnt:   len(items),
			shift: VECTOR_SHIFT,
			root:  EMPTY_PERSISTENT_VECTOR_NODE,
			tail:  items,
		}
	}
	return CreateVector(items)
}

func CreateLazilyPersistentVector(obj interface{}) IPersistentVector {
	switch o := obj.(type) {
	case IReduceInit:
		return CreateVectorFromIReduceInit(o)
	case ISeq:
		return CreateVectorFromISeq(o)
	default:
		return CreateOwningLazilyPersistentVector(RT.ToArray(obj))
	}
}
