package lang

type IPersistentStack interface {
	IPersistentCollection

	Peek() interface{}
	Pop() IPersistentStack
}
