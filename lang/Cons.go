package lang

/*
	Cons

	Implements abstract class ASeq

	Implements Serializable
 */

type Cons struct {
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

/*
	Abstract methods (Cons)
*/

func (c *Cons) Cons(o interface{}) IPersistentCollection {
	return ASeq_Cons(c, o)
}

func (c *Cons) Empty() IPersistentCollection {
	return ASeq_Empty(c)
}

func (c *Cons) Equals(o interface{}) bool {
	return ASeq_Equals(c, o)
}

func (c *Cons) Equiv(o interface{}) bool {
	return ASeq_Equiv(c, o)
}

func (c *Cons) HashCode() int {
	return ASeq_HashCode(c)
}

func (c *Cons) HashEq() int {
	return ASeq_HashEq(c)
}

func (c *Cons) Seq() ISeq {
	return ASeq_Seq(c)
}

func (c *Cons) String() string {
	return ASeq_String(c)
}