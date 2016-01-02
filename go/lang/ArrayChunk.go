package lang

// NOTE: Implements IChunk, Serializable
type ArrayChunk struct {
	array []interface{}
	off   int
	end   int
}

func (a *ArrayChunk) Nth(i int, notFound interface{}) interface{} {
	if i >= 0 && i < a.Count() {
		return a.array[a.off+i]
	}
	return notFound
}

func (a *ArrayChunk) Count() int {
	return a.end - a.off
}

func (a *ArrayChunk) DropFirst() IChunk {
	if a.off == a.end {
		panic("DropFirst of empty chunk")
	}
	return &ArrayChunk{
		array: a.array,
		off:   a.off + 1,
		end:   a.end,
	}
}

func (a *ArrayChunk) Reduce(f IFn, start interface{}) interface{} {
	ret := f.Invoke(start, a.array[a.off])
	if RT.IsReduced(ret) {
		return ret
	}
	for x := a.off + 1; x < a.end; x++ {
		ret = f.Invoke(ret, a.array[x])
		if RT.IsReduced(ret) {
			return ret
		}
	}
	return ret
}
