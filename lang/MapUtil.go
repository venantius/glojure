package lang

/*
	A collection of map-related utilities. These exist to help us deal with the
	fact that we can't make Glojure maps share the same interface as Go maps.
*/

// TODO
func getEntrySetFromMapPrimitive(m map[interface{}]interface{}) IPersistentSet {
	return nil
}

func MapEntrySet(obj interface{}) IPersistentSet {
	switch o := obj.(type) {
	case map[interface{}]interface{}:
		return getEntrySetFromMapPrimitive(o)
	case APersistentMap:
		return APersistentMap_EntrySet(o)
	}
	panic(InvalidTypeException)
}

// Retrieve a map's keys.
func MapKeys(m interface{}) []interface{} {
	mTyped := m.(map[interface{}]interface{})
	keys := make([]interface{}, len(mTyped))

	i := 0
	for k := range mTyped {
		keys[i] = k
		i++
	}
	return keys
}
