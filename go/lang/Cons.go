package lang

// NOTE: Implements Serializable
type Cons struct {
	*ASeq

	_meta  IPersistentMap
	_first interface{}
	_more  ISeq
}

func (c *Cons) First() interface{} {
	return c._first
}

func (c *Cons) Next() ISeq {
	return c._more.Seq()
}

func (c *Cons) More() ISeq {
	if c._more == nil {
		return EMPTY_PERSISTENT_LIST
	}
	return c._more
}

func (c *Cons) Count() int {
	return 1 + RT.Count(c._more)
}

func (c *Cons) WithMeta(meta IPersistentMap) *Cons {
	return &Cons{
		_meta:  meta,
		_first: c._first,
		_more:  c._more,
	}
}
