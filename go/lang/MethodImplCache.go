package lang

type MethodImplCache struct {
	protocol IPersistentMap
	methodk  Keyword
	shift    int
	mask     int
	table    []interface{}
	m        map[interface{}]interface{}
}

type Entry struct {
	fn IFn
	// c  *type TODO (could maybe use reflect.Type?)
}

var mre Entry

// TODO
