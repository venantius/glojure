package lang

type MapEntry struct {
	AMapEntry

	_key interface{}
	_val interface{}
}

func CreateMapEntry(key interface{}, val interface{}) *MapEntry {
	return &MapEntry{
		_key: key,
		_val: val,
	}
}

func (m *MapEntry) Key() interface{} {
	return m._key
}

func (m *MapEntry) Val() interface{} {
	return m._val
}

// TODO: Are these two functions even necessary?
// We don't need to support the Java API.

func (m *MapEntry) GetKey() interface{} {
	return m.Key()
}

func (m *MapEntry) GetValue() interface{} {
	return m.Val()
}