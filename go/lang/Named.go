package lang

type Named interface {
	GetNamespace() string
	GetName() string
}
