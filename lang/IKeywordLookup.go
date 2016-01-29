package lang

type IKeywordLookup interface {
	GetLookupThunk(k Keyword) ILookupThunk
}
