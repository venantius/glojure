package lang

type ILookupSite interface {
	Fault(target interface{}) ILookupThunk
}
