package lang

// TODO: extends Map.Entry from java.util.map
type IMapEntry interface {
	key() interface{}
	val() interface{}
}
