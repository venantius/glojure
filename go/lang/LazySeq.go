package lang

// NOTE: Implements ISeq, Sequential, List, IPending, IHashEq
type LazySeq struct {
	*Obj

	fn IFn
	sv interface{}
	s  ISeq
}

// TODO...quite a bit more here.

// TODO
func (l *LazySeq) Seq() ISeq {
	return nil
}
