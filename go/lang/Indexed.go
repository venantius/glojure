package lang

type Indexed interface {
	Counted

	nth(i int, notFound interface{}) interface{}
}
